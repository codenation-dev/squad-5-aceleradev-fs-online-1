package service

import (
	"app/domain/builder"
	"app/domain/errors"
	"app/domain/model"
	"app/domain/service/engine"
	"app/resources/repository"
	"bufio"
	"io"
	"mime/multipart"
)

// Customer interface
type Customer interface {
	Parse(file multipart.File) (*model.CustomerInsert, error)
	read(reader io.Reader) (*model.CustomerInsert, error)
	createCustomer(customer *model.Customer) (*model.Customer, error)
}

// CustomerService struct
type CustomerService struct {
	Repository repository.CustomerDB
	Alert      engine.EngineAlert
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
		customerDB, err := cs.createCustomer(&customer)
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

func (cs CustomerService) createCustomer(customer *model.Customer) (*model.Customer, error) {

	err := cs.Repository.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	cs.Alert.Customers() <- *customer
	return customer, err
}
