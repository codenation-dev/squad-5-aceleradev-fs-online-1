package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"

	"github.com/go-xorm/xorm"
)

// LoginDB interface
type LoginDB interface {
	Authorization(q validator.Login) error
}

// LoginRepository struct
type LoginRepository struct {
	DB *xorm.Engine
}

// Authorization autoriza um usuario
func (lr LoginRepository) Authorization(q validator.Login) error {

	var user model.User
	t, err := filters(&q, lr.DB).Get(&user)

	if err != nil {
		return err
	}
	if t == false {
		return errors.AuthorizationError
	}
	return nil
}

func filters(q *validator.Login, DB *xorm.Engine) *xorm.Session {

	s := DB.NoCache()

	s = s.Where("Username = ?", q.Username).And("Password = ?", q.Password)

	return s
}
