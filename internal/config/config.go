// Package config loads process settings from the environment.
// It does not open connections or start servers.
package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	envHTTPAddr          = "RUNPHASE_HTTP_ADDR"
	envDatabaseURL       = "RUNPHASE_DATABASE_URL"
	envTemporalAddr      = "RUNPHASE_TEMPORAL_ADDR"
	envTemporalNamespace = "RUNPHASE_TEMPORAL_NAMESPACE"

	defaultHTTPAddr          = ":8081"
	defaultDatabaseURL       = "postgres://runphase:runphase@localhost:5432/runphase?sslmode=disable"
	defaultTemporalAddr      = "localhost:7233"
	defaultTemporalNamespace = "default"
)

// Config is the resolved settings for one Runphase process.
type Config struct {
	HTTP     HTTP
	Database Database
	Temporal Temporal
}

// HTTP is where a future API process would listen.
type HTTP struct {
	Addr string
}

// Database identifies the Runphase product database.
// It is not Temporal's database.
type Database struct {
	URL string
}

// Temporal identifies the local Temporal frontend and namespace.
type Temporal struct {
	Address   string
	Namespace string
}

// Load reads RUNPHASE_* environment variables.
// An unset variable uses the local-development default.
// A variable set to an empty string is an error.
func Load() (Config, error) {
	cfg := Config{
		HTTP:     HTTP{Addr: lookup(envHTTPAddr, defaultHTTPAddr)},
		Database: Database{URL: lookup(envDatabaseURL, defaultDatabaseURL)},
		Temporal: Temporal{
			Address:   lookup(envTemporalAddr, defaultTemporalAddr),
			Namespace: lookup(envTemporalNamespace, defaultTemporalNamespace),
		},
	}
	if err := cfg.validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func lookup(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func (c Config) validate() error {
	var errs []error
	if c.HTTP.Addr == "" {
		errs = append(errs, fmt.Errorf("%s must not be empty", envHTTPAddr))
	}
	if c.Database.URL == "" {
		errs = append(errs, fmt.Errorf("%s must not be empty", envDatabaseURL))
	}
	if c.Temporal.Address == "" {
		errs = append(errs, fmt.Errorf("%s must not be empty", envTemporalAddr))
	}
	if c.Temporal.Namespace == "" {
		errs = append(errs, fmt.Errorf("%s must not be empty", envTemporalNamespace))
	}
	return errors.Join(errs...)
}
