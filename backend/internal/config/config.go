package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	AppEnv string
	HTTP   HTTPConfig
	DB     DBConfig
	Auth   AuthConfig
	Upload UploadConfig
	AI     AIConfig
	Log    LogConfig
}

type HTTPConfig struct {
	Port           string
	CORSOrigins    string
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type AuthConfig struct {
	Secret      string
	TokenTTL    time.Duration
}

type UploadConfig struct {
	MaxSizeMB  int64
	TempDir    string
}

type AIConfig struct {
	TestTimeoutSeconds      int
	ExtractionTimeoutSeconds int
}

type LogConfig struct {
	Level string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	tokenTTL, err := parseDuration("AUTH_TOKEN_TTL_MINUTES", 60)
	if err != nil {
		return nil, fmt.Errorf("config: AUTH_TOKEN_TTL_MINUTES: %w", err)
	}

	maxUploadMB, err := parseInt64("MAX_UPLOAD_SIZE_MB", 20)
	if err != nil {
		return nil, fmt.Errorf("config: MAX_UPLOAD_SIZE_MB: %w", err)
	}

	aiTestTimeout, err := parseInt("AI_PROVIDER_TEST_TIMEOUT_SECONDS", 15)
	if err != nil {
		return nil, fmt.Errorf("config: AI_PROVIDER_TEST_TIMEOUT_SECONDS: %w", err)
	}

	aiExtractTimeout, err := parseInt("AI_EXTRACTION_TIMEOUT_SECONDS", 120)
	if err != nil {
		return nil, fmt.Errorf("config: AI_EXTRACTION_TIMEOUT_SECONDS: %w", err)
	}

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		HTTP: HTTPConfig{
			Port:        getEnv("HTTP_PORT", "8080"),
			CORSOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"),
		},
		DB: DBConfig{
			Host:     getEnv("DATABASE_HOST", "localhost"),
			Port:     getEnv("DATABASE_PORT", "3306"),
			Name:     requireEnv("DATABASE_NAME"),
			User:     requireEnv("DATABASE_USER"),
			Password: requireEnv("DATABASE_PASSWORD"),
		},
		Auth: AuthConfig{
			Secret:   requireEnv("AUTH_SECRET"),
			TokenTTL: time.Duration(tokenTTL) * time.Minute,
		},
		Upload: UploadConfig{
			MaxSizeMB: maxUploadMB,
			TempDir:   getEnv("UPLOAD_TEMP_DIR", "/tmp/finance-ai-app"),
		},
		AI: AIConfig{
			TestTimeoutSeconds:       aiTestTimeout,
			ExtractionTimeoutSeconds: aiExtractTimeout,
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}

	return cfg, nil
}

// DSN returns the MariaDB data source name.
func (c *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		c.User, c.Password, c.Host, c.Port, c.Name,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required environment variable %q is not set", key))
	}
	return v
}

func parseDuration(key string, defaultMinutes int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultMinutes, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("must be an integer, got %q", v)
	}
	return n, nil
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
