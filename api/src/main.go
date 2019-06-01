package main

import (
	"app/application"
	"app/resources/repository"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := repository.NewConnection()
	log.Printf("Server started")

	router := application.NewRouter(db)

	log.Fatal(router.Run(":5000"))
}
