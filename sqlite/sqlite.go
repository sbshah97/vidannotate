package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB is a struct that wraps a *sql.DB connection
type DB struct {
	*sql.DB
}

// New opens a new connection to a SQLite database
func New(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return &DB{db}, nil
}

// Close closes the underlying database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

// Create inserts a new row into the database
func (db *DB) Create(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// Read queries the database and returns the rows
func (db *DB) Read(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

// Update modifies existing rows in the database
func (db *DB) Update(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// Delete removes rows from the database
func (db *DB) Delete(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}
