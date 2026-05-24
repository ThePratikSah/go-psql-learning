package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL        string
	JWTSecret          string
	JWTExpiry          time.Duration
	RefreshTokenExpiry time.Duration
	Port               string
	Env                string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		DatabaseURL:        getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/nextask?sslmode=disable"),
		JWTSecret:          getEnv("JWT_SECRET", "default-secret-change-me"),
		JWTExpiry:          getEnvDuration("JWT_EXPIRY", 15*time.Minute),
		RefreshTokenExpiry: getEnvDuration("REFRESH_TOKEN_EXPIRY", 7*24*time.Hour),
		Port:               getEnv("PORT", "8080"),
		Env:                getEnv("ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	dur, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return dur
}
