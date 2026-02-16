package handler

import "database/sql"

var db *sql.DB

// SetDB sets the database instance used by handlers.
func SetDB(d *sql.DB) {
	db = d
}
