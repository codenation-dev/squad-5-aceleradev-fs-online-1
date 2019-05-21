package repository

import (
	"app/config/database"
	"database/sql"
	"fmt"
)

//NewConnection retorna uma nova conexão do banco de dados
// return: *sql.DB
func NewConnection() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
