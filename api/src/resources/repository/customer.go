package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"strings"

	"github.com/go-xorm/xorm"
)

// CustomerDB interface
type CustomerDB interface {
	CreateCustomer(custumer *model.Customer) error
	Get(custumer *model.Customer) (bool, error)
	UpdateCustomer(customer *model.Customer) error
	ListCustomer(q *validator.CustomerListRequest) (*[]model.Customer, error)
	CountCustomers() (int64, error)
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

// UpdateCustomer atualiza um customer
func (r CustomerRepository) UpdateCustomer(customer *model.Customer) error {

	_, err := r.DB.Id(customer.ID).Update(customer)

	if err != nil {
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			return errors.DuplicatedCustomerError
		}

		return err
	}

	return nil

}

// ListCustomer Lista customer
func (r CustomerRepository) ListCustomer(q *validator.CustomerListRequest) (*[]model.Customer, error) {

	var customers []model.Customer

	if q.Limit == 0 {
		q.Limit = 20
	}

	if err := r.DB.Limit(q.Limit, q.Offset).Find(&customers); err != nil {
		return nil, err
	}
	return &customers, nil
}

// CountCustomers quantidade total de customer
func (r CustomerRepository) CountCustomers() (int64, error) {

	customer := model.Customer{}

	total, err := r.DB.Count(customer)

	if err != nil {
		return 0, err
	}

	return total, nil
}
