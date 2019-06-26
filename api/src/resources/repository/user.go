package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"strings"

	"github.com/go-xorm/xorm"
)

// UserDB interface
type UserDB interface {
	GetUser(id string) (*model.User, error)
	Get(user *model.User) (bool, error)
	CreateUser(user *model.User) error
	ListUser(q *validator.UserListRequest) (*[]model.User, error)
	CountUsers(q *validator.UserListRequest) (int64, error)
	UpdateUser(u *model.User) error
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

// Get Recupera um User
func (r UserRepository) Get(user *model.User) (bool, error) {
	return r.DB.Get(user)
}

// CreateUser cria um novo usuário
func (r UserRepository) CreateUser(user *model.User) error {
	_, err := r.DB.InsertOne(user)
	if err != nil {
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			return errors.DuplicatedUserError
		}
		return err
	}
	return err
}

// ListUser lista os usuários
func (r UserRepository) ListUser(q *validator.UserListRequest) (*[]model.User, error) {
	var users []model.User
	if q.Limit == 0 {
		q.Limit = 20
	}

	if err := addFilters(q, r.DB).Limit(q.Limit, q.Offset).Find(&users); err != nil {
		return nil, err
	}

	return &users, nil
}

// CountUsers conta os usuários
func (r UserRepository) CountUsers(q *validator.UserListRequest) (int64, error) {
	user := new(model.User)
	total, err := addFilters(q, r.DB).Count(user)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// UpdateUser atualiza um usuário já existente
func (r UserRepository) UpdateUser(u *model.User) error {
	_, err := r.DB.Id(u.ID).Update(u)
	if err != nil {
		if strings.Index(strings.ToLower(err.Error()), "unique constraint") >= 0 {
			return errors.DuplicatedUserError
		}

		return err
	}
	return nil
}

func addFilters(q *validator.UserListRequest, DB *xorm.Engine) *xorm.Session {
	s := DB.NoCache()
	if q.Name != "" {
		s = s.Where("Name like ?", "%"+q.Name+"%")
	}
	if q.Email != "" {
		s = s.Where("Email like ?", "%"+q.Email+"%")
	}
	return s
}
