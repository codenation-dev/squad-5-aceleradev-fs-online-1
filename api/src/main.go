package main

import (
	"app/config/db"
	"app/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	helloWorld()
}

func helloWorld() {
	fmt.Println("Hello World")
	repository.Init()
	initDB()
}

func initDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
