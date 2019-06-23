package service

import (
	"app/domain/builder"
	"app/domain/errors"
	"app/domain/model"
	"app/domain/service/engine"
	"app/domain/validator"
	"app/resources/repository"
	"bufio"
	"io"
	"mime/multipart"
)

// Customer interface
type Customer interface {
	Parse(file multipart.File) (*model.CustomerInsert, error)
	read(reader io.Reader) (*model.CustomerInsert, error)
	CreateCustomer(customer *model.Customer) (*model.Customer, error)
	UpdateCustomer(id string, customer *model.Customer) (*model.Customer, error)
	ListCustomer(q *validator.CustomerListRequest) (*model.CustomerList, error)
}

// CustomerService struct
type CustomerService struct {
	Repository repository.CustomerDB
	Alert      engine.Alert
}

// Parse recebe um nome de arquivo e retorna o seu conteudo
func (cs CustomerService) Parse(file multipart.File) (*model.CustomerInsert, error) {

	defer file.Close()
	cl, err := cs.read(file)

	if err != nil {
		return cl, err
	}
	return cl, nil
}

func (cs CustomerService) read(reader io.Reader) (*model.CustomerInsert, error) {

	var customers []model.Customer
	ci := &model.CustomerInsert{}
	r := bufio.NewReader(reader)

	var errs []error
	i := 0
	for ; ; i++ {
		s, _, err := r.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return ci, err
			}

		}

		customer := model.Customer{
			Name: string(s),
			ID:   builder.NewULID(),
		}
		customerDB, err := cs.CreateCustomer(&customer)
		if err != nil {
			errs = append(errs, err)
		} else {
			customers = append(customers, *customerDB)
		}
	}

	ci.Success = len(customers)
	ci.AlreadyExist = len(errs)

	if ci.Success == 0 {
		return ci, errors.AllDuplicatedCustomerError
	}
	return ci, nil
}

// CreateCustomer percistencia do customer
func (cs CustomerService) CreateCustomer(customer *model.Customer) (*model.Customer, error) {

	err := cs.Repository.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	cs.Alert.Customers() <- *customer
	return customer, err
}

// UpdateCustomer ...
func (cs CustomerService) UpdateCustomer(id string, customer *model.Customer) (*model.Customer, error) {

	customer.ID = id
	err := cs.Repository.UpdateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// ListCustomer lista os customer com total
func (cs CustomerService) ListCustomer(q *validator.CustomerListRequest) (*model.CustomerList, error) {

	customerList := model.CustomerList{}

	customers, err := cs.Repository.ListCustomer(q)
	if err != nil {
		return nil, err
	}

	customerList.Data = *customers
	total, err := cs.Repository.CountCustomers()

	if err != nil {
		return nil, err
	}
	customerList.Records = total

	return &customerList, nil
}
