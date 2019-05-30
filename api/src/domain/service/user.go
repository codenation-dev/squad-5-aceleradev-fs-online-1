package service

import (
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
)

// Users interface
type Users interface {
	ListUsers(q *validator.UserListRequest) (*model.UserList, error)
}

// UserService struct
type UserService struct {
	Repository repository.UserDB
}

// ListUsers Lista usuários com total
func (us UserService) ListUsers(q *validator.UserListRequest) (*model.UserList, error) {
	userList := model.UserList{}

	users, err := us.Repository.ListUser(q)
	if err != nil {
		return nil, err
	}
	userList.Data = *users

	total, err := us.Repository.CountUsers(q)
	if err != nil {
		return nil, err
	}
	userList.Records = total

	return &userList, nil
}
