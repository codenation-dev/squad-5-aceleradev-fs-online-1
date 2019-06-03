package service

import (
	"app/domain/errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"app/domain/model"
	"strings"
	"testing"

)

type mockDBCustomer struct {
	customers  []model.Customer
	err   error
	count int64
}

func (mk mockDBCustomer) CreateCustomer(custumer *model.Customer) error{

	for _, c := range mk.customers {
		if c.Name == custumer.Name {
			return errors.DuplicatedCustomerError
		}
	}

	return mk.err
}

func TestCustomerService_read(t *testing.T) {
	
	mock := mockDBCustomer{}
	cs := CustomerService{mock}
	c, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))

	cl := model.CustomerList{
		Records:3,
		Data: []model.Customer{
			{Name:"customer 1"},
		 	{Name:"customer 2"},
			{Name:"customer 3"},
		},

	}
	assert.Nil(t, err)
	assert.Equal(t, c.Records, cl.Records)
	assert.Equal(t, c.Data, cl.Data)
}

func TestCustomerService_read_AllDuplicatedCustomerError(t *testing.T) {
	
	mock := mockDBCustomer{customers:[]model.Customer{
		{Name:"customer 1"},
		{Name:"customer 2"},
		{Name:"customer 3"},
	}}
	cs := CustomerService{mock}
	_, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))
	
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.AllDuplicatedCustomerError)
}

func TestCustomerService_read_ListDuplicatedCustomerError(t *testing.T) {
	
	mock := mockDBCustomer{customers:[]model.Customer{
		{Name:"customer 1"},
		{Name:"customer 2"},
	}}
	cs := CustomerService{mock}
	c, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))
	
	cl := model.CustomerList{
		Records:1,
		Data: []model.Customer{
			{Name:"customer 3"},
		},

	}

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.ListDuplicatedCustomerError)
	assert.Equal(t, c.Records, cl.Records)
	assert.Equal(t, c.Data, cl.Data)
}
