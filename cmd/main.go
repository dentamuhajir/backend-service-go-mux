package main

import (
	"github.com/dentamuhajir/backend-service-go-mysql/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	// err = server.Start()
	// if err != nil {
	//log.Fatal(err)
	//}
}
