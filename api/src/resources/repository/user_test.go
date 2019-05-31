package repository

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func initDB(runMigrations bool) *UserRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if runMigrations {
		RunMigrations(db)
	}

	return &UserRepository{DB: db}
}

func TestUserRepository_GetUser(t *testing.T) {
	uRepo := initDB(true)
	defer uRepo.DB.Close()

	u := model.User{
		ID:       "1111",
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}
	uRepo.DB.InsertOne(u)

	newUser, err := uRepo.GetUser("1111")
	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, u.ID, newUser.ID)
	assert.Equal(t, u.Username, newUser.Username)
	assert.Equal(t, u.Password, newUser.Password)
	assert.Equal(t, u.Name, newUser.Name)
	assert.Equal(t, u.Email, newUser.Email)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)
}

func TestUserRepository_GetUser_NotFound(t *testing.T) {
	uRepo := initDB(true)
	defer uRepo.DB.Close()

	newUser, err := uRepo.GetUser("1111")
	assert.Nil(t, err)
	assert.Nil(t, newUser)
}

func TestUserRepository_CreateUser(t *testing.T) {
	uRepo := initDB(true)
	defer uRepo.DB.Close()

	uc := validator.UserCreation{
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}

	newUser, err := uRepo.CreateUser(&uc)
	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, uc.Username, newUser.Username)
	assert.Equal(t, uc.Password, newUser.Password)
	assert.Equal(t, uc.Name, newUser.Name)
	assert.Equal(t, uc.Email, newUser.Email)
	assert.NotNil(t, newUser.ID)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)

	user := model.User{}
	ok, err := uRepo.DB.ID(newUser.ID).Get(&user)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, ok)
	newUser.CreatedAt = user.CreatedAt
	newUser.UpdatedAt = user.UpdatedAt
	assert.Equal(t, &user, newUser)
}

func TestUserRepository_CreateUser_DuplicatedUser(t *testing.T) {
	uRepo := initDB(true)
	defer uRepo.DB.Close()

	uc := validator.UserCreation{
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}

	newUser, err := uRepo.CreateUser(&uc)
	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, uc.Username, newUser.Username)
	assert.Equal(t, uc.Password, newUser.Password)
	assert.Equal(t, uc.Name, newUser.Name)
	assert.Equal(t, uc.Email, newUser.Email)
	assert.NotNil(t, newUser.ID)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)

	user := model.User{}
	ok, err := uRepo.DB.ID(newUser.ID).Get(&user)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, ok)
	newUser.CreatedAt = user.CreatedAt
	newUser.UpdatedAt = user.UpdatedAt
	assert.Equal(t, &user, newUser)

	newUser, err = uRepo.CreateUser(&uc)
	assert.NotNil(t, err)
	assert.Nil(t, newUser)
	assert.Equal(t, errors.DuplicatedUserError, err)
}