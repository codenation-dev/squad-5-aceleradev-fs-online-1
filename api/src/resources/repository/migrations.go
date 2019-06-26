package repository

import (
	"app/domain/model"
	"log"

	"github.com/go-xorm/xorm"
)

// RunMigrations executa a sincronização com o banco de dados
func RunMigrations(db *xorm.Engine) {
	log.Println("Migrations starting")

	err := db.Sync(
		new(model.User),
		new(model.Customer),
		new(model.PublicAgent),
		new(model.Alert),
		new(model.AlertUser),
	)

	if err != nil {
		log.Println("Migrations error: ", err)
		panic(err)
	}

	log.Println("Migrations success")
}
