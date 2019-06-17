package repository

import (
	"app/domain/model"
	"time"

	"github.com/go-xorm/xorm"
)

// AlertDB interface
type AlertDB interface {
	GetAlert(id string) (*model.Alert, error)
}

// AlertRepository struct
type AlertRepository struct {
	DB *xorm.Engine
}

type alert struct {
	ID          string             `xorm:"varchar(26) pk 'id'"`
	Type        model.AlertType    `xorm:"int notnull"`
	Description string             `xorm:"varchar(100) notnull"`
	Customer    *model.Customer    `xorm:"extends"`
	PublicAgent *model.PublicAgent `xorm:"extends"`
	User        *model.User        `xorm:"extends"`
	CreatedAt   time.Time          `xorm:"notnull created"`
	UpdatedAt   time.Time          `xorm:"notnull updated"`
}

// GetAlert recupera os detalhes do Alerta
func (r AlertRepository) GetAlert(id string) (*model.Alert, error) {
	a := alert{}
	ok, err := r.DB.Table("alert").
		Join("LEFT", "customer", "alert.customer_id = customer.id").
		Join("LEFT", "public_agent", "alert.public_agent_id = public_agent.id").
		Join("LEFT", "\"user\"", "alert.user_id = \"user\".id").
		Where("alert.id = ?", id).
		Get(&a)
	if ok == false || err != nil {
		return nil, err
	}

	//fmt.Printf("ALERT %#v\n", a)

	var users []model.User
	err = r.DB.Table([]string{"user", "u"}).
		Join("INNER", "alert_user", "u.id = alert_user.user_id").
		Where("alert_user.alert_id = ?", id).
		Find(&users)

	if err != nil {
		return nil, err
	}

	//fmt.Printf("USERS %#v\n", users)

	return mapAlert(a, users), nil
}

func mapAlert(a alert, users []model.User) *model.Alert {
	var customer *model.Customer
	if a.Customer != nil && a.Customer.Name != "" {
		customer = a.Customer
	}
	var publicAgent *model.PublicAgent
	if a.PublicAgent != nil && a.PublicAgent.Name != "" {
		publicAgent = a.PublicAgent
	}
	var user *model.User
	if a.User != nil && a.User.Name != "" {
		user = a.User
	}
	return &model.Alert{
		ID:            a.ID,
		Type:          a.Type,
		Description:   a.Description,
		Customer:      customer,
		PublicAgent:   publicAgent,
		User:          user,
		UsersReceived: users,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}
