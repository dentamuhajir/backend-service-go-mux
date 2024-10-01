package database 

import (
    "database/sql"
    "log"

    "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(cfg mysql.Config) (*sql.DB, error) {
    db, err := sql.Open("mysql", cfg.FormatDSN())

    if err != nil {
        return nil, err   
    }

    if err := db.Ping(); err != nil {
        db.Close()  
        return nil, err
    }

    log.Println("Database connection successful:", cfg.FormatDSN())  

    return db, nil
}