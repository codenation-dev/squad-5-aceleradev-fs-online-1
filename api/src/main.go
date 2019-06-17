package main

import (
	"app/application/controller"
	"app/application"
	"app/resources/repository"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := repository.NewConnection()
	controller.InitSendEmail(db)
	log.Printf("Server started")

	router := application.NewRouter(db)

	log.Fatal(router.Run(":5000"))
	
}
