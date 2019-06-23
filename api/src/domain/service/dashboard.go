package service

import (
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
)

// Dasboards interface
type Dasboards interface {
	GetData(q validator.DashboardRequest) (*model.Dashboard, error)
}

// DasboardService struct
type DasboardService struct {
	Repository repository.DasboardDB
}

// GetData recupera os dados estatisticos
func (ds DasboardService) GetData(q validator.DashboardRequest) (*model.Dashboard, error) {
	return ds.Repository.GetData(q)
}
