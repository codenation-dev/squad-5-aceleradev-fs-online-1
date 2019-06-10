package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func initDBCustomer(runMigrations bool) *CustomerRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if runMigrations {
		RunMigrations(db)
	}

	return &CustomerRepository{DB: db}
}

func TestCustomerRepository_CreateCustomer(t *testing.T) {
	uRepo := initDBCustomer(true)
	defer uRepo.DB.Close()

	newCustomer := model.Customer{
		ID:   "1111",
		Name: "test",
	}

	err := uRepo.CreateCustomer(&newCustomer)
	assert.Nil(t, err)

	customer := model.Customer{}
	ok, err := uRepo.DB.ID(newCustomer.ID).Get(&customer)
	assert.Nil(t, err)
	assert.NotNil(t, ok)
	assert.True(t, ok)
	newCustomer.CreatedAt = customer.CreatedAt
	newCustomer.UpdatedAt = customer.UpdatedAt
	assert.Equal(t, customer, newCustomer)
}

func TestCustomerRepository_CreateCustomer_DuplicatedCustomer(t *testing.T) {
	uRepo := initDBCustomer(true)
	defer uRepo.DB.Close()

	newCustomer := model.Customer{
		ID:   "1111",
		Name: "test",
	}

	err := uRepo.CreateCustomer(&newCustomer)
	assert.Nil(t, err)

	customer := model.Customer{}
	ok, err := uRepo.DB.ID(newCustomer.ID).Get(&customer)
	assert.Nil(t, err)
	assert.NotNil(t, ok)
	assert.True(t, ok)
	newCustomer.CreatedAt = customer.CreatedAt
	newCustomer.UpdatedAt = customer.UpdatedAt
	assert.Equal(t, customer, newCustomer)

	err = uRepo.CreateCustomer(&newCustomer)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.DuplicatedCustomerError)
}
