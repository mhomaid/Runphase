package config

import (
	"os"
	"strings"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	unset(t, envHTTPAddr)
	unset(t, envDatabaseURL)
	unset(t, envTemporalAddr)
	unset(t, envTemporalNamespace)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.HTTP.Addr != defaultHTTPAddr {
		t.Errorf("HTTP.Addr = %q, want %q", cfg.HTTP.Addr, defaultHTTPAddr)
	}
	if cfg.Database.URL != defaultDatabaseURL {
		t.Errorf("Database.URL = %q, want %q", cfg.Database.URL, defaultDatabaseURL)
	}
	if cfg.Temporal.Address != defaultTemporalAddr {
		t.Errorf("Temporal.Address = %q, want %q", cfg.Temporal.Address, defaultTemporalAddr)
	}
	if cfg.Temporal.Namespace != defaultTemporalNamespace {
		t.Errorf("Temporal.Namespace = %q, want %q", cfg.Temporal.Namespace, defaultTemporalNamespace)
	}
}

func TestLoadOverrides(t *testing.T) {
	t.Setenv(envHTTPAddr, ":9090")
	t.Setenv(envDatabaseURL, "postgres://example:example@db:5432/example?sslmode=disable")
	t.Setenv(envTemporalAddr, "temporal:7233")
	t.Setenv(envTemporalNamespace, "runphase")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.HTTP.Addr != ":9090" {
		t.Errorf("HTTP.Addr = %q", cfg.HTTP.Addr)
	}
	if cfg.Database.URL != "postgres://example:example@db:5432/example?sslmode=disable" {
		t.Errorf("Database.URL = %q", cfg.Database.URL)
	}
	if cfg.Temporal.Address != "temporal:7233" {
		t.Errorf("Temporal.Address = %q", cfg.Temporal.Address)
	}
	if cfg.Temporal.Namespace != "runphase" {
		t.Errorf("Temporal.Namespace = %q", cfg.Temporal.Namespace)
	}
}

func TestLoadRejectsEmptyOverrides(t *testing.T) {
	cases := []string{
		envHTTPAddr,
		envDatabaseURL,
		envTemporalAddr,
		envTemporalNamespace,
	}
	for _, key := range cases {
		t.Run(key, func(t *testing.T) {
			unset(t, envHTTPAddr)
			unset(t, envDatabaseURL)
			unset(t, envTemporalAddr)
			unset(t, envTemporalNamespace)
			t.Setenv(key, "")

			_, err := Load()
			if err == nil {
				t.Fatal("Load() error = nil, want an error")
			}
			if !strings.Contains(err.Error(), key+" must not be empty") {
				t.Fatalf("Load() error = %q, want it to name %s", err, key)
			}
		})
	}
}

func unset(t *testing.T, key string) {
	t.Helper()
	previous, ok := os.LookupEnv(key)
	if err := os.Unsetenv(key); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if !ok {
			_ = os.Unsetenv(key)
			return
		}
		_ = os.Setenv(key, previous)
	})
}
