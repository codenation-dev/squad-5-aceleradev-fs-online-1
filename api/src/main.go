package main

import (
	"app/repository"
	_ "github.com/lib/pq"
)

func main() {

	start()
}

func start() {
	repository.NewConnection()
}

