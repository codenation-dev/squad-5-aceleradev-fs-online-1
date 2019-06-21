package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"strings"

	"github.com/go-xorm/xorm"
)

// CustomerDB interface
type CustomerDB interface {
	CreateCustomer(custumer *model.Customer) error
	Get(custumer *model.Customer) (bool, error)
}

// CustomerRepository struct
type CustomerRepository struct {
	DB *xorm.Engine
}

// CreateCustomer cria um novo customer
func (r CustomerRepository) CreateCustomer(custumer *model.Customer) error {

	_, err := r.DB.Insert(custumer)
	if err != nil {
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			return errors.DuplicatedCustomerError
		}
		return err
	}
	return nil
}

// Get Recupera um cliente
func (r CustomerRepository) Get(custumer *model.Customer) (bool, error) {
	return r.DB.Get(custumer)
}
