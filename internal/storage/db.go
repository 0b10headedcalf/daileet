package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"modernc.org/sqlite"
)

func init() {
	sql.Register("sqlite3", &sqlite.Driver{})
}

// OpenDB opens the SQLite database at ./daileet.db relative to the project root.
// If the DB doesn't exist, it creates it and runs migrations + seeds the Blind 75.
func OpenDB() (*sql.DB, error) {
	// Try to find project root by looking for go.mod
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dbPath := filepath.Join(cwd, "daileet.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	schema := `
CREATE TABLE IF NOT EXISTS problems (
    id            INTEGER PRIMARY KEY,
    title         TEXT NOT NULL,
    title_slug    TEXT UNIQUE NOT NULL,
    difficulty    TEXT,
    pattern       TEXT,
    url           TEXT,
    interval      REAL DEFAULT 0,
    repetitions   INTEGER DEFAULT 0,
    ease_factor   REAL DEFAULT 2.5,
    due_date      DATETIME,
    last_reviewed DATETIME,
    status        TEXT DEFAULT 'new'
);

CREATE TABLE IF NOT EXISTS config (
    key   TEXT PRIMARY KEY,
    value TEXT
);
`
	_, err := db.Exec(schema)
	return err
}

func isSeeded(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM problems").Scan(&count)
	return count > 0, err
}
