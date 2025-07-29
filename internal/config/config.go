package config

import (
	"os"
	"strings"
)

type Config struct {
	PostgresUrl string
}

func Load() Config {
	return Config{
		PostgresUrl: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/meuapp?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return strings.TrimSpace(val)
	}

	return fallback
}
