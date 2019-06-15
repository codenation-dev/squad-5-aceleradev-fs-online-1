package repository

import (
	"app/domain/model"
	"strings"

	"github.com/go-xorm/xorm"
)

// PublicAgentDB interface
type PublicAgentDB interface {
	CreateOrUpdatePublicAgent(publicAgent *model.PublicAgent) error
}

// PublicAgentRepository struct
type PublicAgentRepository struct {
	DB *xorm.Engine
}

// CreateOrUpdatePublicAgent cria ou atualiza o funcionário publico
func (r PublicAgentRepository) CreateOrUpdatePublicAgent(publicAgent *model.PublicAgent) error {
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
				return err
			}

			uPublicAgent := model.PublicAgent{Checked: publicAgent.Checked}
			if ePublicAgent.Salary != publicAgent.Salary {
				uPublicAgent.UpdatedAt = publicAgent.Checked
				uPublicAgent.Salary = publicAgent.Salary
			}
			_, err = r.DB.Id(ePublicAgent.ID).Update(uPublicAgent)
		}
	}

	return err
}
