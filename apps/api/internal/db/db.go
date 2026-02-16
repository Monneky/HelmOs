package db

import (
	"database/sql"
	"path/filepath"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	once sync.Once
	db   *sql.DB
)

// Open opens the SQLite database under dataDir (creates dir if needed).
func Open(dataDir string) (*sql.DB, error) {
	var err error
	once.Do(func() {
		dbPath := filepath.Join(dataDir, "helmos.db")
		db, err = sql.Open("sqlite", dbPath)
		if err != nil {
			return
		}
		_, _ = db.Exec("PRAGMA journal_mode=WAL")
	})
	return db, err
}

// DB returns the global DB instance (must call Open first).
func DB() *sql.DB {
	return db
}

// Close closes the global DB.
func Close() error {
	if db == nil {
		return nil
	}
	return db.Close()
}
