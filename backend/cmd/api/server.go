package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"finance-ai-app/backend/internal/config"
	"finance-ai-app/backend/internal/database"
	router "finance-ai-app/backend/internal/http"
)

const (
	readTimeout       = 15 * time.Second
	readHeaderTimeout = 5 * time.Second
	writeTimeout      = 120 * time.Second
	idleTimeout       = 60 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	db, err := database.Connect(cfg.DB.DSN(), database.Config{
		MaxOpenConns:    cfg.DB.MaxOpenConns,
		MaxIdleConns:    cfg.DB.MaxIdleConns,
		ConnMaxLifetime: cfg.DB.ConnMaxLifetime,
		ConnMaxIdleTime: cfg.DB.ConnMaxIdleTime,
		PingTimeout:     cfg.DB.PingTimeout,
	})
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	defer db.Close()

	log.Printf(

		"connected to database %s@%s:%s/%s",
		cfg.DB.User,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	r := router.New(cfg.HTTP.CORSOrigins)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.HTTP.Port),
		Handler:           r,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	return serve(server, cfg.AppEnv)
}

// start runs the HTTP server and sends unexpected server errors to serverErrors.
func start(server *http.Server, appEnv string, serverErrors chan<- error) {
	log.Printf("starting server on %s (env=%s)", server.Addr, appEnv)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		serverErrors <- err
	}
}

func serve(server *http.Server, appEnv string) error {
	// Channel to receive unexpected server errors.
	serverErrors := make(chan error, 1)
	go start(server, appEnv, serverErrors)

	// Channel to listen for interrupt or terminate signals from the OS.
	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdownSignals)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdownSignals:
		log.Printf("received signal %s, shutting down server", sig)

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("graceful shutdown: %w", err)
		}

		log.Println("server stopped gracefully")
		return nil
	}
}
