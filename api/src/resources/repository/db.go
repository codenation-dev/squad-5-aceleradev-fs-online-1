package repository

import (
	"app/application/config/database"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
)

//NewConnection retorna uma nova conexão do banco de dados
// return: *xorm.Engine
func NewConnection() *xorm.Engine {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.DBname)

	db, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.ShowSQL(false)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")

	RunMigrations(db)

	return db
}
