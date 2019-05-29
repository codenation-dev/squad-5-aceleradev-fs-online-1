package repository

import (
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRunMigrations(t *testing.T) {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	RunMigrations(db)

	tables, err := db.DBMetas()

	assert.Nil(t, err)
	assert.Len(t, tables, 1)
}

func TestRunMigrations_Error(t *testing.T) {
	db := xorm.Engine{}

	assert.Panics(t, func() {
		RunMigrations(&db)
	})
}
