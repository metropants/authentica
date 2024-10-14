package main

import (
	"github.com/metropants/authentica/internal/config"
	"github.com/metropants/authentica/internal/database"
	"log/slog"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("an error occurred loading config", "error", err)
		return
	}

	pool, err := database.New(cfg)
	if err != nil {
		slog.Error("an error occurred connection to the database", "error", err)
		return
	}
	defer pool.Close()

	if cfg.AutoMigrate {
		err := database.Migrate(pool)
		if err != nil {
			slog.Error("an error occurred migrating the database", "error", err)
			return
		}
	}
}
