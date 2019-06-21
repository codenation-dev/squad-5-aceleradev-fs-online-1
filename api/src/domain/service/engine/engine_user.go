package engine

import (
	"app/domain/model"
	"log"
)

func (eas EngineAlertService) proccessUsers() {
	for user := range users {
		eas.checkUser(user)
	}
}

func (eas EngineAlertService) checkUser(user model.User) {
	customer := model.Customer{
		Name: user.Name,
	}
	has, err := eas.CustomerDB.Get(&customer)
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
