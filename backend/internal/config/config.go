package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	AppEnv   string
	HTTP     HTTPConfig
	DB       DBConfig
	Auth     AuthConfig
	Security SecurityConfig
	Upload   UploadConfig
	AI       AIConfig
	Log      LogConfig
}

type HTTPConfig struct {
	Port        string
	CORSOrigins []string
}

type DBConfig struct {
	Host            string
	Port            string
	Name            string
	User            string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	PingTimeout     time.Duration
}

type AuthConfig struct {
	Secret   string
	TokenTTL time.Duration
}

type SecurityConfig struct {
	EncryptionKey string
}

type UploadConfig struct {
	MaxSizeMB int64
	TempDir   string
}

type AIConfig struct {
	TestTimeout       time.Duration
	ExtractionTimeout time.Duration
}

type LogConfig struct {
	Level string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	// Load .env in development-style workflows. If the file is missing,
	// keep going and rely on OS environment variables.
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("config: .env load failed: %w", err)
	}

	databaseName, err := requireEnv("DATABASE_NAME")
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_NAME: %w", err)
	}

	databaseUser, err := requireEnv("DATABASE_USER")
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_USER: %w", err)
	}

	databasePassword, err := requireEnv("DATABASE_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_PASSWORD: %w", err)
	}

	authSecret, err := requireEnv("AUTH_SECRET")
	if err != nil {
		return nil, fmt.Errorf("config: AUTH_SECRET: %w", err)
	}

	encryptionKey, err := requireEnv("ENCRYPTION_KEY")
	if err != nil {
		return nil, fmt.Errorf("config: ENCRYPTION_KEY: %w", err)
	}

	tokenTTLMinutes, err := parseInt("AUTH_TOKEN_TTL_MINUTES", 1440)
	if err != nil {
		return nil, fmt.Errorf("config: AUTH_TOKEN_TTL_MINUTES: %w", err)
	}

	maxUploadMB, err := parseInt64("MAX_UPLOAD_SIZE_MB", 20)
	if err != nil {
		return nil, fmt.Errorf("config: MAX_UPLOAD_SIZE_MB: %w", err)
	}

	aiTestTimeoutSeconds, err := parseInt("AI_PROVIDER_TEST_TIMEOUT_SECONDS", 60)
	if err != nil {
		return nil, fmt.Errorf("config: AI_PROVIDER_TEST_TIMEOUT_SECONDS: %w", err)
	}

	aiExtractionTimeoutSeconds, err := parseInt("AI_EXTRACTION_TIMEOUT_SECONDS", 120)
	if err != nil {
		return nil, fmt.Errorf("config: AI_EXTRACTION_TIMEOUT_SECONDS: %w", err)
	}

	dbMaxOpenConns, err := parseInt("DATABASE_MAX_OPEN_CONNS", 25)
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_MAX_OPEN_CONNS: %w", err)
	}

	dbMaxIdleConns, err := parseInt("DATABASE_MAX_IDLE_CONNS", 5)
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_MAX_IDLE_CONNS: %w", err)
	}

	dbConnMaxLifetimeMinutes, err := parseInt("DATABASE_CONN_MAX_LIFETIME_MINUTES", 5)
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_CONN_MAX_LIFETIME_MINUTES: %w", err)
	}

	dbConnMaxIdleTimeMinutes, err := parseInt("DATABASE_CONN_MAX_IDLE_TIME_MINUTES", 1)
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_CONN_MAX_IDLE_TIME_MINUTES: %w", err)
	}

	dbPingTimeoutSeconds, err := parseInt("DATABASE_PING_TIMEOUT_SECONDS", 5)
	if err != nil {
		return nil, fmt.Errorf("config: DATABASE_PING_TIMEOUT_SECONDS: %w", err)
	}

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		HTTP: HTTPConfig{
			Port:        getEnv("HTTP_PORT", "8080"),
			CORSOrigins: parseCSV(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173")),
		},
		DB: DBConfig{
			Host:            getEnv("DATABASE_HOST", "localhost"),
			Port:            getEnv("DATABASE_PORT", "3306"),
			Name:            databaseName,
			User:            databaseUser,
			Password:        databasePassword,
			MaxOpenConns:    dbMaxOpenConns,
			MaxIdleConns:    dbMaxIdleConns,
			ConnMaxLifetime: time.Duration(dbConnMaxLifetimeMinutes) * time.Minute,
			ConnMaxIdleTime: time.Duration(dbConnMaxIdleTimeMinutes) * time.Minute,
			PingTimeout:     time.Duration(dbPingTimeoutSeconds) * time.Second,
		},
		Auth: AuthConfig{
			Secret:   authSecret,
			TokenTTL: time.Duration(tokenTTLMinutes) * time.Minute,
		},
		Security: SecurityConfig{
			EncryptionKey: encryptionKey,
		},
		Upload: UploadConfig{
			MaxSizeMB: maxUploadMB,
			TempDir:   getEnv("UPLOAD_TEMP_DIR", "/tmp/finance-ai-app"),
		},
		AI: AIConfig{
			TestTimeout:       time.Duration(aiTestTimeoutSeconds) * time.Second,
			ExtractionTimeout: time.Duration(aiExtractionTimeoutSeconds) * time.Second,
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}

	return cfg, nil
}

// DSN returns the MariaDB data source name.
func (c DBConfig) DSN() string {
	cfg := mysql.Config{
		User:                 c.User,
		Passwd:               c.Password,
		Net:                  "tcp",
		Addr:                 c.Host + ":" + c.Port,
		DBName:               c.Name,
		ParseTime:            true,
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
	}

	return cfg.FormatDSN()
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func requireEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("required environment variable is not set")
	}

	return v, nil
}

func parseInt(key string, defaultVal int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal, nil
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("must be an integer, got %q", v)
	}

	return n, nil
}

func parseInt64(key string, defaultVal int64) (int64, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal, nil
	}

	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("must be an integer, got %q", v)
	}

	return n, nil
}

func parseCSV(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
