package main

import (
	"fmt"
	"log"
	"net/http"

	"finance-ai-app/backend/internal/config"
	"finance-ai-app/backend/internal/database"
	router "finance-ai-app/backend/internal/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg.DB.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("connected to database %s@%s:%s/%s", cfg.DB.User, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	r := router.New(cfg.HTTP.CORSOrigins)

	addr := fmt.Sprintf(":%s", cfg.HTTP.Port)
	log.Printf("starting server on %s (env=%s)", addr, cfg.AppEnv)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
