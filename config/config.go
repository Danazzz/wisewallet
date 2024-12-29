package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return &Config{
		DatabaseURL: databaseURL,
	}
}