package builder

import (
	"app/domain/model"
	"app/domain/validator"
)

//CustomerCreationToCustomer ...
func CustomerCreationToCustomer(customerCreate *validator.CustomerCreation) *model.Customer {

	return &model.Customer{
		ID:     NewULID(),
		Name:   customerCreate.Name,
		Salary: customerCreate.Salary,
	}
}
