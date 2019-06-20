package builder

import (
	"app/domain/model"
)

// AlertBuilder cria um novo alerta
func AlertBuilder(t model.AlertType, c *model.Customer, p *model.PublicAgent, u *model.User, to []model.User) *model.Alert {
	return &model.Alert{
		ID:            NewULID(),
		Type:          t,
		Description:   getAlertDescription(t),
		Customer:      c,
		PublicAgent:   p,
		User:          u,
		UsersReceived: to,
	}
}

func getAlertDescription(t model.AlertType) string {
	switch t {
	case model.PublicAgentType:
		return "Cliente é funcionário publico"
	case model.BiggerSalaryType:
		return "Valor do salário em destaque"
	case model.BankEmployeeType:
		return "Cliente é funcionário do banco"
	default:
		return ""
	}
}
