package service

import (
	"app/domain/builder"
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
	"log"
)

// EngineAlert interface
type EngineAlert interface {
	Init()
	PublicAgents() chan model.PublicAgent
}

// EngineAlertService struct
type EngineAlertService struct {
	CustomerDB repository.CustomerDB
	AlertDB    repository.AlertDB
	UserDB     repository.UserDB
}

var publicAgents chan model.PublicAgent

// Init inicializar o EngineAlert
func (eas EngineAlertService) Init() {
	publicAgents = make(chan model.PublicAgent)

	go eas.proccessPublicAgents()
}

// PublicAgents retorna o canal onde receberemos as dados modificados
func (eas EngineAlertService) PublicAgents() chan model.PublicAgent {
	return publicAgents
}

func (eas EngineAlertService) proccessPublicAgents() {
	for publicAgent := range publicAgents {
		eas.checkCustomer(publicAgent)
	}
}

func (eas EngineAlertService) checkCustomer(publicAgent model.PublicAgent) {
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
	}
}

func (eas EngineAlertService) getUsers() (*[]model.User, error) {
	q := &validator.UserListRequest{Limit: -1}

	return eas.UserDB.ListUser(q)
}

func (eas EngineAlertService) createAlert(t model.AlertType, c *model.Customer, p *model.PublicAgent, u *model.User) error {

	users, err := eas.getUsers()
	if err != nil {
		log.Println("Database Error: ", err)
		return err
	}

	a := builder.AlertBuilder(model.PublicAgentType, c, p, nil, *users)

	if err := eas.AlertDB.CreateAlert(a); err != nil {
		log.Println("CreateAlert Error: ", err)
		return err
	}

	return nil
}
