package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() *Config {
	return &Config{
		ServerAddress: getEnv("PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
