package repository

import (
	"app/domain/builder"
	"app/domain/model"
	"strings"

	"github.com/go-xorm/xorm"
)

// PublicAgentDB interface
type PublicAgentDB interface {
	CreateOrUpdatePublicAgent(publicAgent *model.PublicAgent) (bool, error)
}

// PublicAgentRepository struct
type PublicAgentRepository struct {
	DB *xorm.Engine
}

// CreateOrUpdatePublicAgent cria ou atualiza o funcionário publico
func (r PublicAgentRepository) CreateOrUpdatePublicAgent(publicAgent *model.PublicAgent) (bool, error) {
	var err error
	_, err = r.DB.InsertOne(publicAgent)
	if err != nil {
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			var ePublicAgent = model.PublicAgent{
				Name:       publicAgent.Name,
				Occupation: publicAgent.Occupation,
				Department: publicAgent.Department,
			}
			_, err = r.DB.Get(&ePublicAgent)
			if err != nil {
				return false, err
			}

			publicAgent.ID = ePublicAgent.ID
			updated := false
			uPublicAgent := model.PublicAgent{Checked: publicAgent.Checked}
			if builder.Round(ePublicAgent.Salary) != builder.Round(publicAgent.Salary) {
				uPublicAgent.UpdatedAt = publicAgent.Checked
				uPublicAgent.Salary = publicAgent.Salary
				updated = true
			}
			_, err = r.DB.Id(ePublicAgent.ID).Update(uPublicAgent)
			return updated, err
		}

		return false, err
	}

	return true, nil
}
