package service

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDBCustomer struct {
	customers []model.Customer
	err       error
	count     int64
}

func (mk mockDBCustomer) CreateCustomer(custumer *model.Customer) error {

	for _, c := range mk.customers {
		if c.Name == custumer.Name {
			return errors.DuplicatedCustomerError
		}
	}

	return mk.err
}

func (mk mockDBCustomer) Get(custumer *model.Customer) (bool, error) {
	return true, mk.err
}

func (mk mockDBCustomer) UpdateCustomer(customer *model.Customer) error {
	return mk.err
}

func (mk mockDBCustomer) ListCustomer(q *validator.CustomerListRequest) (*[]model.Customer, error) {

	return &[]model.Customer{}, nil
}

func (mk mockDBCustomer) CountCustomers() (int64, error) {

	return int64(0), nil
}

var (
	c chan model.Customer
	p chan model.PublicAgent
	u chan model.User
)

type mockEngineAlert struct {
}

func (mea mockEngineAlert) Init() {
	c = make(chan model.Customer)
	p = make(chan model.PublicAgent)
	u = make(chan model.User)

	go func() {
		for d := range c {
			fmt.Println(d)
		}
		for d := range p {
			fmt.Println(d)
		}
		for d := range u {
			fmt.Println(d)
		}
	}()
}

func (mea mockEngineAlert) PublicAgents() chan model.PublicAgent {
	return p
}
func (mea mockEngineAlert) Customers() chan model.Customer {
	return c
}
func (mea mockEngineAlert) Users() chan model.User {
	return u
}

func TestCustomerService_read(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

	mock := mockDBCustomer{}
	cs := CustomerService{
		Repository: mock,
		Alert:      alert,
	}
	c, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))

	ci := model.CustomerInsert{
		Success: 3,
	}

	assert.Nil(t, err)
	assert.Equal(t, c.Success, ci.Success)
}

func TestCustomerService_read_AllDuplicatedCustomerError(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

	mock := mockDBCustomer{customers: []model.Customer{
		{Name: "customer 1"},
		{Name: "customer 2"},
		{Name: "customer 3"},
	}}
	cs := CustomerService{
		Repository: mock,
		Alert:      alert,
	}

	c, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))

	ci := model.CustomerInsert{
		AlreadyExist: 3,
	}
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.AllDuplicatedCustomerError)
	assert.Equal(t, c.AlreadyExist, ci.AlreadyExist)
	assert.Equal(t, c.Success, ci.Success)
}

func TestCustomerService_read_ListDuplicatedCustomerError(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

	mock := mockDBCustomer{customers: []model.Customer{
		{Name: "customer 1"},
	}}
	cs := CustomerService{
		Repository: mock,
		Alert:      alert,
	}
	c, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))

	ci := model.CustomerInsert{
		Success:      2,
		AlreadyExist: 1,
	}

	assert.Nil(t, err)
	assert.Equal(t, c.Success, ci.Success)
	assert.Equal(t, c.AlreadyExist, ci.AlreadyExist)

}

func TestCustomerService_UpdateCustomer(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

	mock := mockDBCustomer{err: nil}

	cs := CustomerService{
		Repository: mock,
		Alert:      alert,
	}

	customer := &model.Customer{
		ID:     "1111",
		Name:   "test",
		Salary: 100.10,
	}

	c, err := cs.UpdateCustomer("1111", customer)

	assert.Nil(t, err)
	assert.Equal(t, c, customer)
}

func TestCustomerService_UpdateCustomer_DuplicatedCustomerError(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

	mock := mockDBCustomer{err: errors.DuplicatedCustomerError}

	cs := CustomerService{
		Repository: mock,
		Alert:      alert,
	}

	customer := &model.Customer{
		ID:     "1111",
		Name:   "test",
		Salary: 100.10,
	}

	_, err := cs.UpdateCustomer("1111", customer)

	assert.NotNil(t, err)
	assert.Equal(t, err, errors.DuplicatedCustomerError)
}
