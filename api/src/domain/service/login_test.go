package service

import (
	"app/domain/errors"
	"app/domain/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockDBLogin struct
type mockDBLogin struct {
	err error
}

func (mk mockDBLogin) Authorization(q validator.Login) error {
	return mk.err
}

func TestLoginService_Authorization(t *testing.T) {

	mock := mockDBLogin{
		err: nil,
	}

	ls := LoginService{mock}

	q := validator.Login{
		Username: "test",
		Password: "test",
	}

	token, err := ls.Authorization(q)

	assert.Nil(t, err)
	assert.NotZero(t, len(token.Token))
}

func TestLoginService_Authorization_AuthorizationError(t *testing.T) {

	mock := mockDBLogin{
		err: errors.AuthorizationError,
	}

	ls := LoginService{mock}

	q := validator.Login{
		Username: "test",
		Password: "test",
	}

	token, err := ls.Authorization(q)

	assert.NotNil(t, err)
	assert.Nil(t, token)
}
