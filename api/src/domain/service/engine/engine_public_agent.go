package engine

import (
	"app/domain/model"
	"log"
)

func (eas EngineAlertService) proccessPublicAgents() {
	for publicAgent := range publicAgents {
		eas.checkPublicAgent(publicAgent)
	}
}

func (eas EngineAlertService) checkPublicAgentSalary(customer model.Customer, publicAgent model.PublicAgent) {
	if publicAgent.Salary >= BiggerSalaryAlert {
		if err := eas.createAlert(model.BiggerSalaryType, &customer, &publicAgent, nil); err != nil {
			return
		}
	}
}

func (eas EngineAlertService) checkPublicAgent(publicAgent model.PublicAgent) {
	customer := model.Customer{
		Name: publicAgent.Name,
	}
	has, err := eas.CustomerDB.Get(&customer)
	if err != nil {
		log.Println("Database Error: ", err)
		return
	}

	if has {
		if err = eas.createAlert(model.PublicAgentType, &customer, &publicAgent, nil); err != nil {
			return
		}

		eas.checkPublicAgentSalary(customer, publicAgent)
	}
}
