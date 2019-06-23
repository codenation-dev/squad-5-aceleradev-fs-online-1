package repository

import (
	"app/domain/model"
	"app/domain/validator"
	"app/domain/errors"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/stretchr/testify/assert"
)

func initDBLogin(runMigrations bool) *LoginRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if runMigrations {
		RunMigrations(db)
	}

	return &LoginRepository{DB: db}
}
func TestLoginRepository_Authorization(t *testing.T) {

	uRepoLogin := initDBLogin(false)
	defer uRepoLogin.DB.Close()
	uRepoUser := initDBUser(true)
	defer uRepoUser.DB.Close()

	newUser := model.User{
		ID:       "1111",
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}

	err := uRepoUser.CreateUser(&newUser)
	assert.Nil(t, err)

	q := validator.Login{
		Username: "test",
		Password: "test",
	}

	err = uRepoLogin.Authorization(q)
	assert.Nil(t, err)
}

func TestLoginRepository_Authorization_AuthorizationError(t *testing.T) {

	uRepoLogin := initDBLogin(false)
	defer uRepoLogin.DB.Close()
	uRepoUser := initDBUser(true)
	defer uRepoUser.DB.Close()

	newUser := model.User{
		ID:       "1111",
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}

	err := uRepoUser.CreateUser(&newUser)
	assert.Nil(t, err)

	q := validator.Login{
		Username: "test",
		Password: "test2",
	}

	err = uRepoLogin.Authorization(q)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.AuthorizationError)
}
