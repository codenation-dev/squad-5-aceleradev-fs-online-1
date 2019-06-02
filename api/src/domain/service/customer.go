package service

import (
	"app/domain/errors"
	"app/domain/model"
	"app/resources/repository"
	"bufio"
	"io"
	"mime/multipart"
)

// Customer interface
type Customer interface {
	Parse(fileName string) (*model.CustomerList, error)
	read(reader io.Reader) (*model.CustomerList, error)
	createCustomer(customer *model.Customer) (*model.Customer, error)
}

// CustomerService struct
type CustomerService struct {
	Repository repository.CustomerDB
}

// Parse recebe um nome de arquivo e retorna o seu conteudo
func (cs CustomerService) Parse(file multipart.File) (*model.CustomerList, error) {

	defer file.Close()

	cl, err := cs.read(file)

	if err != nil {
		return cl, err
	}
	return cl, nil
}

func (cs CustomerService) read(reader io.Reader) (*model.CustomerList, error) {

	var customers []model.Customer
	cl := &model.CustomerList{}
	r := bufio.NewReader(reader)

	var errs []error
	i := 0
	for ; ; i++ {
		s, _, err := r.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return cl, err
			}

		}

		if len(s) == 0 {
			continue
		} else {

			customer := &model.Customer{Name: string(s)}

			customer, err = cs.createCustomer(customer)
	
			if err != nil {
				if err == errors.DuplicatedCustomerError {
					errs = append(errs, err)
				}
			}
			customers = append(customers, *customer)
		}


	}

	cl = &model.CustomerList{
		Data:    customers,
		Records: int64(len(customers)),
	}

	if len(errs) == i {
		return nil, errors.AllDuplicatedCustomerError
	} else if len(errs) != i && len(errs) != 0 {
		return cl, errors.ListDuplicatedCustomerError
	}

	return cl, nil
}

func (cs CustomerService) createCustomer(customer *model.Customer) (*model.Customer, error) {

	err := cs.Repository.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return customer, err
}
