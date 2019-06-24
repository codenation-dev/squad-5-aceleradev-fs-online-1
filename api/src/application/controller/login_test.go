package controller

import (
	"app/domain/model"
	"testing"
	"app/domain/validator"
	"app/domain/errors"
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockLogin struct {
	token *model.Token
	err error
}

func (mk mockLogin) Authorization(q validator.Login) (*model.Token, error){

	return mk.token, mk.err
}


func TestLoginController_Authorization(t *testing.T) {

	mock := mockLogin{
		token: &model.Token{
			Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M",
		},
		err: nil,
	}

	lc := LoginController{mock}

	router := gin.Default()

	router.POST("/auth", lc.Authorization)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password": "123456"
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M\"}", w.Body.String())

}

func TestLoginController_Authorization_ValidationErrorLoginPass(t *testing.T) {

	mock := mockLogin{
		token: &model.Token{
			Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M",
		},
		err: nil,
	}

	lc := LoginController{mock}

	router := gin.Default()

	router.POST("/auth", lc.Authorization)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password": "test"
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Password\",\"message\":\"Field validation for 'Password' failed on the 'min' tag\"}]", w.Body.String())

}

func TestLoginController_Authorization_ValidationErrorLoginName(t *testing.T) {

	mock := mockLogin{
		token: &model.Token{
			Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M",
		},
		err: nil,
	}

	lc := LoginController{mock}

	router := gin.Default()

	router.POST("/auth", lc.Authorization)

	b := bytes.NewReader([]byte(`{
		"username": "",
		"password": "123456"
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Username\",\"message\":\"Field validation for 'Username' failed on the 'required' tag\"}]", w.Body.String())

}

func TestLoginController_Authorization_AuthorizationError(t *testing.T) {

	mock := mockLogin{
		token: nil,
		err: errors.AuthorizationError,
	}

	lc := LoginController{mock}

	router := gin.Default()

	router.POST("/auth", lc.Authorization)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password": "123456"
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "[{\"message\":\"Invalid username or password\"}]", w.Body.String())

}
