package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DatabaseURI string
	AutoMigrate bool
}

func Load() (*Config, error) {
	migrate, err := strconv.ParseBool(os.Getenv("DATABASE_AUTO_MIGRATE"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DatabaseURI: os.Getenv("DATABASE_URI"),
		AutoMigrate: migrate,
	}
	return cfg, nil
}
