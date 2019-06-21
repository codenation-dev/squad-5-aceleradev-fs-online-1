package service

import (
	"app/domain/errors"
	"app/domain/model"
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

var (
	c chan model.Customer
	p chan model.PublicAgent
	u chan model.User
)

func init() {
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

type mockEngineAlert struct {
}

func (mea mockEngineAlert) Init() {
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
	go func() {
		c, err := cs.read(strings.NewReader(fmt.Sprintf("customer 1\ncustomer 2\ncustomer 3")))

		ci := model.CustomerInsert{
			Success: 3,
		}

		assert.Nil(t, err)
		assert.Equal(t, c.Success, ci.Success)
	}()

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
