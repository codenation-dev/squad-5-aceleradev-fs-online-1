package engine

import (
	"app/domain/model"
	"log"
)

func (eas AlertService) proccessCustomers() {
	for customer := range customers {
		eas.checkCustomer(customer)
	}
}

func (eas AlertService) checkCustomer(customer model.Customer) {
	eas.checkCustomerSalary(customer)

	user := model.User{
		Name: customer.Name,
	}
	has, err := eas.UserDB.Get(&user)
	if err != nil {
		log.Println("Database Error: ", err)
		return
	}

	if has {
		if err = eas.createAlert(model.BankEmployeeType, &customer, nil, &user); err != nil {
			return
		}
	}
}

func (eas AlertService) checkCustomerSalary(customer model.Customer) {
	if float64(customer.Salary) >= BiggerSalaryAlert {
		if err := eas.createAlert(model.BiggerSalaryType, &customer, nil, nil); err != nil {
			return
		}
	}
}
