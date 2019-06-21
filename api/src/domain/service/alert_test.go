package service

import (
	"app/domain/model"
	"app/domain/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDBAlert struct {
	alert *model.Alert
	err   error
	list  []model.AlertItem
	count int64
}

func (m mockDBAlert) GetAlert(id string) (*model.Alert, error) {
	return m.alert, m.err
}

func (m mockDBAlert) ListAlerts(q *validator.AlertListRequest) ([]model.AlertItem, error) {
	return m.list, m.err
}

func (m mockDBAlert) CountAlerts(q *validator.AlertListRequest) (int64, error) {
	return m.count, m.err
}

func TestAlertService_GetAlert(t *testing.T) {
	alert := model.Alert{
		ID:            "12345678901234567890123456",
		Type:          model.PublicAgentType,
		Description:   "test",
		Customer:      nil,
		PublicAgent:   nil,
		User:          nil,
		UsersReceived: []model.User{},
	}
	mock := mockDBAlert{
		alert: &alert,
		err:   nil,
	}
	as := AlertService{mock}

	a, err := as.GetAlert("12345678901234567890123456")
	assert.Nil(t, err)
	assert.Equal(t, &alert, a)
}
