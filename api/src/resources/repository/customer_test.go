package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
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

func TestCustomerRepository_UpdateCustomer(t *testing.T) {
	uRepo := initDBCustomer(true)
	defer uRepo.DB.Close()

	newCustomer := model.Customer{
		ID:   "1111",
		Name: "test",
	}

	err := uRepo.CreateCustomer(&newCustomer)
	assert.Nil(t, err)

	newCustomer.Name = "test2"
	err = uRepo.UpdateCustomer(&newCustomer)

	var customer model.Customer

	ok, err := uRepo.DB.ID(newCustomer.ID).Get(&customer)
	assert.Nil(t, err)
	assert.NotNil(t, ok)
	assert.True(t, ok)
	assert.Equal(t, customer, model.Customer{
		ID:        "1111",
		Name:      "test2",
		Salary:    0,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	})
}

func TestCustomerRepository_UpdateCustomer_DuplicatedCustomerError(t *testing.T) {
	uRepo := initDBCustomer(true)
	defer uRepo.DB.Close()

	newCustomer := model.Customer{
		ID:   "1111",
		Name: "test",
	}

	newCustomer2 := model.Customer{
		ID:   "2222",
		Name: "test2",
	}

	err := uRepo.CreateCustomer(&newCustomer)
	assert.Nil(t, err)
	err = uRepo.CreateCustomer(&newCustomer2)
	assert.Nil(t, err)

	newCustomer2.Name = "test"
	err = uRepo.UpdateCustomer(&newCustomer2)

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.DuplicatedCustomerError)

}

func TestCustomerRepository_ListCustomer(t *testing.T) {

	uRepo := initDBCustomer(true)
	defer uRepo.DB.Close()

	newCustomer := model.Customer{
		ID:   "1111",
		Name: "test",
	}

	err := uRepo.CreateCustomer(&newCustomer)
	assert.Nil(t, err)

	var q validator.CustomerListRequest
	customers, err := uRepo.ListCustomer(&q)

	assert.Nil(t, err)

	var customer model.Customer
	ok, err := uRepo.DB.ID(newCustomer.ID).Get(&customer)
	assert.Nil(t, err)
	assert.NotNil(t, ok)
	assert.True(t, ok)
	assert.Equal(t, customers, &[]model.Customer{customer})

}

func TestCustomerRepository_CountCustomers(t *testing.T) {

	uRepo := initDBCustomer(true)
	defer uRepo.DB.Close()

	newCustomer := model.Customer{
		ID:   "1111",
		Name: "test",
	}

	newCustomer2 := model.Customer{
		ID:   "2222",
		Name: "test2",
	}

	err := uRepo.CreateCustomer(&newCustomer)
	assert.Nil(t, err)
	err = uRepo.CreateCustomer(&newCustomer2)
	assert.Nil(t, err)

	total, err := uRepo.CountCustomers()

	assert.Nil(t, err)
	assert.Equal(t, total, int64(2))

}
