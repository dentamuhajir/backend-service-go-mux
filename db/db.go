package db

import (
	"database/sql"
	"log"

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

	log.Println("\nDatabase successfully connected => ", cfg.FormatDSN())

	// migrate -database "mysql://root@tcp(127.0.0.1:3306)/godb" -path db/migrations up
	return db, err
}
