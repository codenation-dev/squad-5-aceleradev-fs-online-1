package controller

import (
	"app/domain/model"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockAlert struct {
	alert *model.Alert
	err   error
}

func (ac mockAlert) GetAlert(id string) (*model.Alert, error) {
	return ac.alert, ac.err
}

func TestGetAlert(t *testing.T) {
	objCreated := model.Alert{
		ID:            "12345678901234567890123456",
		Type:          model.PublicAgentType,
		Description:   "test",
		Customer:      nil,
		PublicAgent:   nil,
		User:          nil,
		UsersReceived: []model.User{},
	}
	mock := mockAlert{
		alert: &objCreated,
		err:   nil,
	}
	ac := AlertController{mock}

	router := gin.Default()

	router.GET("/alerts/:id", ac.GetAlert)
	id := "12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var objGet map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &objGet)
	assert.Nil(t, err)
	assert.Equal(t, objGet["id"], id)

	assert.Equal(t, objCreated.Type.String(), objGet["type"])
	assert.Equal(t, objCreated.Description, objGet["description"])
	assert.Nil(t, objGet["customer"])
	assert.Nil(t, objGet["publicAgent"])
	assert.Nil(t, objGet["user"])
	assert.Nil(t, objGet["users_received"])
}

func TestGetAlert_NotFound(t *testing.T) {
	mock := mockAlert{
		alert: nil,
		err:   nil,
	}
	ac := AlertController{mock}

	router := gin.Default()

	router.GET("/alerts/:id", ac.GetAlert)
	id := "12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetAlert_ValidationError(t *testing.T) {
	mock := mockAlert{
		alert: nil,
		err:   nil,
	}
	ac := AlertController{mock}

	router := gin.Default()

	router.GET("/alerts/:id", ac.GetAlert)
	id := "212121"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"ID\",\"message\":\"Field validation for 'ID' failed on the 'len' tag\"}]", w.Body.String())
}

func TestGetAlert_GenericError(t *testing.T) {
	mock := mockAlert{
		alert: nil,
		err:   errors.New("generic error"),
	}
	ac := AlertController{mock}

	router := gin.Default()

	router.GET("/alerts/:id", ac.GetAlert)
	id := "12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "", w.Body.String())
}
