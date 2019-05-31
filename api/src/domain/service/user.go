package service

import (
	"app/domain/builder"
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
)

// Users interface
type Users interface {
	ListUsers(q *validator.UserListRequest) (*model.UserList, error)
	UpdateUser(id string, userCreation *validator.UserCreation) (*model.User, error)
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

// UpdateUser Atualiza um usuário
func (us UserService) UpdateUser(id string, userCreation *validator.UserCreation) (*model.User, error) {
	user := builder.UserCreationToUser(userCreation)
	user.ID = id

	err := us.Repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
