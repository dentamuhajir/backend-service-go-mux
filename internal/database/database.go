package database  // Clear and descriptive name for the package

import (
    "database/sql"
    "log"

    "github.com/go-sql-driver/mysql"
)

// Public function for better visibility
func NewMySQLConnection(cfg mysql.Config) (*sql.DB, error) {
    db, err := sql.Open("mysql", cfg.FormatDSN())

    if err != nil {
        return nil, err // Return the error for caller handling
    }

    if err := db.Ping(); err != nil {
        db.Close() // Close connection on Ping failure
        return nil, err
    }

    log.Println("Database connection successful:", cfg.FormatDSN()) // Informative message

    return db, nil
}