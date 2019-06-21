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

func (m mockDBUser) Get(user *model.User) (bool, error) {
	return true, m.err
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
	alert := mockEngineAlert{}
	alert.Init()

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
	us := UserService{
		Repository: mock,
		Alert:      alert,
	}

	q := validator.UserListRequest{}

	list, err := us.ListUsers(&q)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), list.Records)
	assert.Len(t, list.Data, 1)
}

func TestUserService_ListUsers_ErrorList(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

	mock := mockDBUser{
		user:  nil,
		err:   errors.New("generic error"),
		count: 1,
	}
	us := UserService{
		Repository: mock,
		Alert:      alert,
	}

	q := validator.UserListRequest{}

	list, err := us.ListUsers(&q)
	assert.NotNil(t, err)
	assert.Nil(t, list)
}

func TestUserService_UpdateUser(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

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
	us := UserService{
		Repository: mock,
		Alert:      alert,
	}

	uc := validator.UserCreation{
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}

	go func() {
		u, err := us.UpdateUser(user.ID, &uc)
		assert.Nil(t, err)
		assert.Equal(t, u, &user)
	}()
}

func TestUserService_CreateUser(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

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
	us := UserService{
		Repository: mock,
		Alert:      alert,
	}

	uc := validator.UserCreation{
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}

	go func() {
		u, err := us.CreateUser(&uc)
		assert.Nil(t, err)
		user.ID = u.ID
		assert.Equal(t, u, &user)
	}()
}

func TestUserService_GetUser(t *testing.T) {
	alert := mockEngineAlert{}
	alert.Init()

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
	us := UserService{
		Repository: mock,
		Alert:      alert,
	}

	u, err := us.GetUser("12345678901234567890123456")
	assert.Nil(t, err)
	assert.Equal(t, &user, u)
}
