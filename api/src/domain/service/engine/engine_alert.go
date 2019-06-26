package engine

import (
	"app/domain/builder"
	"app/domain/model"
	"app/domain/validator"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

// BiggerSalaryAlert define o valor do salário que vai gerar o alerta
const (
	BiggerSalaryAlert float64       = 20000
	DefaultExpiration time.Duration = 24 * time.Hour
)

var ch *cache.Cache

func init() {
	ch = cache.New(DefaultExpiration, 10*time.Minute)
}

func (eas AlertService) getUsers() (*[]model.User, error) {
	q := &validator.UserListRequest{Limit: -1}

	return eas.UserDB.ListUser(q)
}

func (eas AlertService) createAlert(t model.AlertType, c *model.Customer, p *model.PublicAgent, u *model.User) error {
	id := builder.GetUniqueID(t, *c)
	_, has := ch.Get(id)
	if has {
		return nil
	}

	ch.Set(id, "ok", DefaultExpiration)

	users, err := eas.getUsers()
	if err != nil {
		log.Println("Database Error: ", err)
		return err
	}

	a := builder.AlertBuilder(t, c, p, u, *users)

	if err := eas.AlertDB.CreateAlert(a); err != nil {
		log.Println("CreateAlert Error: ", err)
		return err
	}

	eas.EmailChannel <- model.Email{
		Recipients: builder.GetRecipients(*users),
		Subject:    a.Description,
		Body:       builder.GetBody(*a),
	}

	return nil
}
