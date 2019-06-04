package repository

import (
	"app/domain/model"
	"log"

	"github.com/go-xorm/xorm"
)

// RunMigrations executa a sincronização com o banco de dados
func RunMigrations(db *xorm.Engine) {
	log.Println("Migrations starting")

	err := db.Sync(new(model.User))
	if err != nil {
		log.Println("Migrations error: ", err)
		panic(err)
	}

	err = db.Sync(new(model.Customer))

	if err != nil {
		log.Println("Migrations error: ", err)
		panic(err)
	}
}
