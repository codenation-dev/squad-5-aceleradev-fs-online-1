package service

import (
	"app/domain/model"
	"app/resources/repository"
)

// Alerts interface
type Alerts interface {
	GetAlert(id string) (*model.Alert, error)
}

// AlertService struct
type AlertService struct {
	Repository repository.AlertDB
}

// GetAlert returna um alerta
func (as AlertService) GetAlert(id string) (*model.Alert, error) {
	return as.Repository.GetAlert(id)
}
