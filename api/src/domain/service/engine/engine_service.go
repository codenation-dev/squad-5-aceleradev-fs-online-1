package engine

import (
	"app/domain/model"
	"app/resources/repository"
)

// EngineAlert interface
type EngineAlert interface {
	Init()
	PublicAgents() chan model.PublicAgent
	Customers() chan model.Customer
	Users() chan model.User
}

// EngineAlertService struct
type EngineAlertService struct {
	CustomerDB   repository.CustomerDB
	AlertDB      repository.AlertDB
	UserDB       repository.UserDB
	EmailChannel chan model.Email
}

var publicAgents chan model.PublicAgent
var customers chan model.Customer
var users chan model.User

// Init inicializar o EngineAlert
func (eas EngineAlertService) Init() {
	publicAgents = make(chan model.PublicAgent)
	customers = make(chan model.Customer)
	users = make(chan model.User)

	go eas.proccessPublicAgents()
	go eas.proccessCustomers()
	go eas.proccessUsers()
}

// PublicAgents retorna o canal onde receberemos as funcionário publicos modificados
func (eas EngineAlertService) PublicAgents() chan model.PublicAgent {
	return publicAgents
}

// Customers retorna o canal onde recebemos os clientes modificados
func (eas EngineAlertService) Customers() chan model.Customer {
	return customers
}

// Users retorna o canal onde recebemos os usuários/funcionários do banco modificados
func (eas EngineAlertService) Users() chan model.User {
	return users
}
