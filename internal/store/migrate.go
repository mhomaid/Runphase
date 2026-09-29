// Package store applies versioned SQL migrations to the Runphase database.
// It does not open Temporal's databases and it is not used by the API server.
package store

import (
	"embed"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	pgxmigrate "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

const migrationsTable = "schema_migrations"

// Outcome is the migration state after a successful run.
type Outcome struct {
	Version        uint
	Dirty          bool
	AlreadyCurrent bool
}

// Up applies pending migrations to the product database.
// ErrNoChange is success: the database is already current.
func Up(databaseURL string) (Outcome, error) {
	cfg, err := pgx.ParseConfig(databaseURL)
	if err != nil || cfg.Database == "" {
		return Outcome{}, errors.New("database url is invalid")
	}
	if cfg.Database == "temporal" || cfg.Database == "temporal_visibility" {
		return Outcome{}, fmt.Errorf("refusing to migrate database %q", cfg.Database)
	}

	source, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return Outcome{}, fmt.Errorf("read migrations: %w", err)
	}

	db := stdlib.OpenDB(*cfg)
	driver, err := pgxmigrate.WithInstance(db, &pgxmigrate.Config{
		MigrationsTable: migrationsTable,
	})
	if err != nil {
		_ = db.Close()
		return Outcome{}, sanitize(err, databaseURL)
	}

	m, err := migrate.NewWithInstance("iofs", source, cfg.Database, driver)
	if err != nil {
		_ = driver.Close()
		return Outcome{}, sanitize(err, databaseURL)
	}

	upErr := m.Up()
	version, dirty, versionErr := m.Version()
	sourceErr, databaseErr := m.Close()

	if upErr != nil && !errors.Is(upErr, migrate.ErrNoChange) {
		return Outcome{}, sanitize(upErr, databaseURL)
	}
	if versionErr != nil {
		return Outcome{}, sanitize(versionErr, databaseURL)
	}
	if dirty {
		return Outcome{}, fmt.Errorf("database migration is dirty at version %d", version)
	}
	if sourceErr != nil || databaseErr != nil {
		return Outcome{}, sanitize(errors.Join(sourceErr, databaseErr), databaseURL)
	}

	return Outcome{
		Version:        version,
		Dirty:          dirty,
		AlreadyCurrent: errors.Is(upErr, migrate.ErrNoChange),
	}, nil
}

func sanitize(err error, databaseURL string) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	parsed, parseErr := url.Parse(databaseURL)
	if parseErr == nil {
		msg = strings.ReplaceAll(msg, databaseURL, "postgres://redacted")
		if parsed.User != nil {
			if password, ok := parsed.User.Password(); ok && password != "" {
				msg = strings.ReplaceAll(msg, password, "redacted")
			}
			msg = strings.ReplaceAll(msg, parsed.User.String(), "redacted")
		}
	}
	return fmt.Errorf("database migration: %s", msg)
}
