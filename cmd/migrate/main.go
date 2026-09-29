// Command migrate applies versioned SQL to the Runphase product database.
package main

import (
	"log/slog"
	"os"

	"github.com/mhomaid/runphase/internal/config"
	"github.com/mhomaid/runphase/internal/store"
)

func main() {
	if err := run(); err != nil {
		slog.Error("database migration failed", "err", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	slog.Info("database migration starting")
	outcome, err := store.Up(cfg.Database.URL)
	if err != nil {
		return err
	}
	if outcome.AlreadyCurrent {
		slog.Info("database already current", "version", outcome.Version, "dirty", outcome.Dirty)
		return nil
	}
	slog.Info("database migration complete", "version", outcome.Version, "dirty", outcome.Dirty)
	return nil
}
