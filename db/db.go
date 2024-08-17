package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

// private is lowercase , public is uppercase
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err // Return the error to be handled by the caller
	}

	if err := db.Ping(); err != nil {
		db.Close() // Ensure the connection is closed if Ping fails
		return nil, err
	}

	return db, err
}
