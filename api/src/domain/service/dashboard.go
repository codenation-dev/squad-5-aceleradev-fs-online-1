package service

import (
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
)

// Dasboards interface
type Dasboards interface {
	GetData(q validator.DashboardRequest) (*model.Dashboard, error)
	ListCustomers(q validator.DashboardCustomerRequest) (*model.DashboardCustomerList, error)
}

// DasboardService struct
type DasboardService struct {
	Repository      repository.DasboardDB
	AlertRepository repository.AlertRepository
}

// GetData recupera os dados estatisticos
func (ds DasboardService) GetData(q validator.DashboardRequest) (*model.Dashboard, error) {
	return ds.Repository.GetData(q)
}

// ListCustomers lista os ultimos alertas
func (ds DasboardService) ListCustomers(q validator.DashboardCustomerRequest) (*model.DashboardCustomerList, error) {
	dashboardCustomerList := model.DashboardCustomerList{}

	customers, err := ds.Repository.ListCustomers(q)
	if err != nil {
		return nil, err
	}
	dashboardCustomerList.Data = customers

	total, err := ds.AlertRepository.CountAlerts(&validator.AlertListRequest{})
	if err != nil {
		return nil, err
	}
	dashboardCustomerList.Records = total

	return &dashboardCustomerList, nil
}
