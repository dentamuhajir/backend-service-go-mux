package main

import (
	"github.com/dentamuhajir/backend-service-go-mysql/cmd/api"
	"github.com/dentamuhajir/backend-service-go-mysql/config"
	"github.com/dentamuhajir/backend-service-go-mysql/db"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	dbase, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server := api.NewAPIServer(":8080", dbase)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	// err = server.Start()
	// if err != nil {
	//log.Fatal(err)
	//}
}
