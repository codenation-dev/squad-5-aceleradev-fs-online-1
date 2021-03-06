package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

// AlertDB interface
type AlertDB interface {
	GetAlert(id string) (*model.Alert, error)
	ListAlerts(q *validator.AlertListRequest) ([]model.AlertItem, error)
	CountAlerts(q *validator.AlertListRequest) (int64, error)
	CreateAlert(a *model.Alert) error
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

	var users []model.User
	err = r.DB.Table([]string{"user", "u"}).
		Join("INNER", "alert_user", "u.id = alert_user.user_id").
		Where("alert_user.alert_id = ?", id).
		Find(&users)

	if err != nil {
		return nil, err
	}

	return mapAlert(a, users), nil
}

// ListAlerts List os alertas
func (r AlertRepository) ListAlerts(q *validator.AlertListRequest) ([]model.AlertItem, error) {
	var alerts []model.AlertItem
	if q.Limit == 0 {
		q.Limit = 20
	}

	if err := addAlertFilters(q, r.DB).
		Select("a.id, a.type, c.name customer_name,a.created_at").
		Limit(q.Limit, q.Offset).
		Find(&alerts); err != nil {

		return nil, err
	}

	return alerts, nil
}

// CountAlerts conta a quantidade de registros
func (r AlertRepository) CountAlerts(q *validator.AlertListRequest) (int64, error) {
	total, err := addAlertFilters(q, r.DB).Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

// CreateAlert salva o alerta no banco
func (r AlertRepository) CreateAlert(a *model.Alert) error {
	_, err := r.DB.InsertOne(a)
	if err != nil {
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			return errors.DuplicatedUserError
		}
		return err
	}
	_, err = r.DB.Insert(getAlertUsers(a))
	if err != nil {
		return err
	}
	return err
}

func getAlertUsers(a *model.Alert) []model.AlertUser {
	s := len(a.UsersReceived)
	users := make([]model.AlertUser, s, s)

	for i, user := range a.UsersReceived {
		users[i] = model.AlertUser{
			User:  user,
			Alert: *a,
		}
	}

	return users
}

func addAlertFilters(q *validator.AlertListRequest, DB *xorm.Engine) *xorm.Session {
	s := DB.Table([]string{"alert", "a"}).
		Join("LEFT", []string{"customer", "c"}, "a.customer_id = c.id").
		Join("LEFT", []string{"\"user\"", "u"}, "a.user_id = u.id")

	if q.Name != "" {
		s = s.Where("u.name like ?", "%"+q.Name+"%")
	}
	if q.Email != "" {
		s = s.Where("u.email like ?", "%"+q.Email+"%")
	}
	if q.Customer != "" {
		s = s.Where("c.name like ?", "%"+q.Customer+"%")
	}
	if q.Type != "" {
		a, _ := model.AlertTypeFromString(q.Type)
		s = s.Where("a.type = ?", *a)
	}
	if !q.DateStart.IsZero() {
		s = s.Where("a.created_at >= ?", q.DateStart)
	}
	if !q.DateFinish.IsZero() {
		s = s.Where("a.created_at <= ?", q.DateFinish)
	}
	return s
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
