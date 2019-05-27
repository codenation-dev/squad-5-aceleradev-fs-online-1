package repository

import (
	"app/domain/builder"
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"log"

	"github.com/go-xorm/xorm"
)

// UserDB interface
type UserDB interface {
	GetUser(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}

// UserRepository struct
type UserRepository struct {
	DB *xorm.Engine
}

// GetUser recupera um usuário por id
func (r UserRepository) GetUser(id string) (*model.User, error) {
	user := model.User{}
	ok, err := r.DB.ID(id).Get(&user)
	if ok == false || err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser cria um novo usuário
func (r UserRepository) CreateUser(userCreation *validator.UserCreation) (*model.User, error) {
	user := builder.UserCreationToUser(userCreation)
	_, err := r.DB.InsertOne(user)
	if err != nil {
		log.Println("CreateUser error", err)
		return nil, errors.DuplicatedUserError
	}
	return user, err
}
