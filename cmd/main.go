package main

import (
	"log"

	"github.com/dentamuhajir/backend-service-go-mysql/cmd/api"
	"github.com/dentamuhajir/backend-service-go-mysql/config"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/database"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := database.NewMySQLConnection(
		mysql.Config{
			User:                 config.Envs.DBUser,
			Passwd:               config.Envs.DBPassword,
			Addr:                 config.Envs.DBAddress,
			DBName:               config.Envs.DBName,
			Net:                  "tcp",
			AllowNativePasswords: true,
			ParseTime:            true,
		},
	)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server := api.NewAPIServer(":8080", database)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	// err = server.Start()
	// if err != nil {
	//log.Fatal(err)
	//}
}
