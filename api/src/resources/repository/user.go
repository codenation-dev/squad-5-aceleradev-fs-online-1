package repository

import (
	"app/domain/builder"
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"strings"

	"github.com/go-xorm/xorm"
)

// UserDB interface
type UserDB interface {
	GetUser(id string) (*model.User, error)
	CreateUser(userCreation *validator.UserCreation) (*model.User, error)
	ListUser(q *validator.UserListRequest) (*[]model.User, error)
	CountUsers(q *validator.UserListRequest) (int64, error)
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
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			return nil, errors.DuplicatedUserError
		}
		return nil, err
	}
	return user, err
}

// ListUser lista os usuários
func (r UserRepository) ListUser(q *validator.UserListRequest) (*[]model.User, error) {
	var users []model.User

	if err := r.DB.Find(&users); err != nil {
		return nil, err
	}

	return &users, nil
}

// CountUsers conta os usuários
func (r UserRepository) CountUsers(q *validator.UserListRequest) (int64, error) {
	user := new(model.User)
	total, err := r.DB.Count(user)
	if err != nil {
		return 0, err
	}

	return total, nil
}
