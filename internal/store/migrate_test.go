package store

import (
	"errors"
	"io/fs"
	"strings"
	"testing"
)

func TestUpRejectsTemporalDatabases(t *testing.T) {
	cases := []string{
		"postgres://runphase:secret@localhost:5432/temporal?sslmode=disable",
		"postgres://runphase:secret@localhost:5432/temporal_visibility?sslmode=disable",
	}
	for _, databaseURL := range cases {
		t.Run(databaseURL, func(t *testing.T) {
			_, err := Up(databaseURL)
			if err == nil {
				t.Fatal("Up() error = nil")
			}
			if strings.Contains(err.Error(), "secret") {
				t.Fatalf("error leaked a credential: %v", err)
			}
		})
	}
}

func TestUpRejectsInvalidURLWithoutLeakingPassword(t *testing.T) {
	_, err := Up("not a url secret-password")
	if err == nil {
		t.Fatal("Up() error = nil")
	}
	if strings.Contains(err.Error(), "secret-password") {
		t.Fatalf("error leaked a credential: %v", err)
	}
}

func TestSanitizeRemovesCredentials(t *testing.T) {
	databaseURL := "postgres://runphase:s3cret@localhost:5432/runphase?sslmode=disable"
	err := sanitize(errors.New("dial "+databaseURL+" failed for user runphase:s3cret"), databaseURL)
	if err == nil {
		t.Fatal("sanitize() error = nil")
	}
	if strings.Contains(err.Error(), "s3cret") || strings.Contains(err.Error(), databaseURL) {
		t.Fatalf("sanitize() = %v", err)
	}
}

func TestBootstrapMigrationHasNoProductTables(t *testing.T) {
	entries, err := fs.ReadDir(migrationFiles, "migrations")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 {
		t.Fatalf("migration files = %d, want 2", len(entries))
	}
	for _, entry := range entries {
		body, err := fs.ReadFile(migrationFiles, "migrations/"+entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		upper := strings.ToUpper(string(body))
		if strings.Contains(upper, "CREATE TABLE") {
			t.Fatalf("%s creates a table:\n%s", entry.Name(), body)
		}
	}
}
