package service

import (
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
)

// Alerts interface
type Alerts interface {
	GetAlert(id string) (*model.Alert, error)
	ListAlerts(q *validator.AlertListRequest) (*model.AlertList, error)
}

// AlertService struct
type AlertService struct {
	Repository repository.AlertDB
}

// GetAlert returna um alerta
func (as AlertService) GetAlert(id string) (*model.Alert, error) {
	return as.Repository.GetAlert(id)
}
