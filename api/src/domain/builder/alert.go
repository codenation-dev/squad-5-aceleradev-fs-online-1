package builder

import (
	"app/domain/model"
	"bytes"
	"fmt"
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

// GetUniqueID recupera o ID identificador do alerta
func GetUniqueID(t model.AlertType, c model.Customer) string {
	return t.String() + c.ID
}

// GetBody recupera o body do e-mail de alerta
func GetBody(a model.Alert) string {
	b := bytes.NewBufferString(a.Description)
	if a.Customer != nil {
		b.WriteString(fmt.Sprintf("\n Cliente: %s", a.Customer.Name))
	}
	if a.User != nil {
		b.WriteString(fmt.Sprintf("\n Funcionário do Banco: %s", a.User.Name))
	}
	if a.PublicAgent != nil {
		b.WriteString(fmt.Sprintf("\n Funcionário publico: %s", a.PublicAgent.Name))
	}
	return b.String()
}
