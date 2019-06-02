package service

import (
	"app/domain/model"
	"app/domain/validator"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDBUser struct {
	user  *model.User
	err   error
	count int64
}

func (m mockDBUser) GetUser(id string) (*model.User, error) {
	return m.user, m.err
}
func (m mockDBUser) CreateUser(user *model.User) error {
	return m.err
}
func (m mockDBUser) ListUser(q *validator.UserListRequest) (*[]model.User, error) {
	if m.user != nil {
		return &[]model.User{*m.user}, m.err
	}
	return nil, m.err
}
func (m mockDBUser) CountUsers(q *validator.UserListRequest) (int64, error) {
	return m.count, m.err
}
func (m mockDBUser) UpdateUser(u *model.User) error {
	return m.err
}

func TestUserService_ListUsers(t *testing.T) {
	mock := mockDBUser{
		user: &model.User{
			ID:       "12345678901234567890123456",
			Username: "test",
			Password: "testsenha",
			Name:     "test nome",
			Email:    "test@mail.com",
		},
		err:   nil,
		count: 1,
	}
	us := UserService{mock}

	q := validator.UserListRequest{}

	list, err := us.ListUsers(&q)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), list.Records)
	assert.Len(t, list.Data, 1)
}

func TestUserService_ListUsers_ErrorList(t *testing.T) {
	mock := mockDBUser{
		user:  nil,
		err:   errors.New("generic error"),
		count: 1,
	}
	us := UserService{mock}

	q := validator.UserListRequest{}

	list, err := us.ListUsers(&q)
	assert.NotNil(t, err)
	assert.Nil(t, list)
}

func TestUserService_UpdateUser(t *testing.T) {
	user := model.User{
		ID:       "12345678901234567890123456",
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}
	mock := mockDBUser{
		user:  &user,
		err:   nil,
		count: 1,
	}
	us := UserService{mock}

	uc := validator.UserCreation{
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}

	u, err := us.UpdateUser(user.ID, &uc)
	assert.Nil(t, err)
	assert.Equal(t, u, &user)
}

func TestUserService_CreateUser(t *testing.T) {
	user := model.User{
		ID:       "12345678901234567890123456",
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}
	mock := mockDBUser{
		user:  &user,
		err:   nil,
		count: 1,
	}
	us := UserService{mock}

	uc := validator.UserCreation{
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}

	u, err := us.CreateUser(&uc)
	assert.Nil(t, err)
	user.ID = u.ID
	assert.Equal(t, u, &user)
}

func TestUserService_GetUser(t *testing.T) {
	user := model.User{
		ID:       "12345678901234567890123456",
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}
	mock := mockDBUser{
		user:  &user,
		err:   nil,
		count: 1,
	}
	us := UserService{mock}

	u, err := us.GetUser("12345678901234567890123456")
	assert.Nil(t, err)
	assert.Equal(t, &user, u)
}
