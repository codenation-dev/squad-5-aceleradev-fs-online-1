package controller

import (
	apierrors "app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type mockUser struct {
	user  *model.User
	err   error
	count int64
}

func (uc mockUser) CreateUser(userCreation *validator.UserCreation) (*model.User, error) {
	return uc.user, uc.err
}
func (uc mockUser) GetUser(id string) (*model.User, error) {
	return uc.user, uc.err
}
func (uc mockUser) ListUsers(q *validator.UserListRequest) (*model.UserList, error) {
	list := model.UserList{
		Records: uc.count,
		Data:    []model.User{*uc.user},
	}
	return &list, uc.err
}
func (uc mockUser) UpdateUser(id string, userCreation *validator.UserCreation) (*model.User, error) {
	return uc.user, uc.err
}

func TestCreateUser(t *testing.T) {
	mock := mockUser{
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
	uc := UserController{mock}

	router := gin.Default()

	router.POST("/users", uc.CreateUser)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var objGet map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &objGet)
	assert.Nil(t, err)
	assert.NotNil(t, objGet["id"])
	assert.Len(t, objGet["id"], 26)
	assert.NotNil(t, objGet["name"])
	assert.NotNil(t, objGet["username"])
	assert.Equal(t, "", objGet["password"])
	assert.NotNil(t, objGet["email"])
}

func TestCreateUser_ValidationErrorUsername(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.POST("/users", uc.CreateUser)

	b := bytes.NewReader([]byte(`{
		"username": "testtesttesttesttesttesttesttesttesttesttest",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Username\",\"message\":\"Field validation for 'Username' failed on the 'max' tag\"}]", w.Body.String())
}

func TestCreateUser_DBError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   errors.New("generic error"),
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.POST("/users", uc.CreateUser)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestCreateUser_DuplicatedUserError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   apierrors.DuplicatedUserError,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.POST("/users", uc.CreateUser)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "[{\"message\":\"Usuário já existe\"}]", w.Body.String())
}

func TestGetUser(t *testing.T) {
	objCreated := model.User{
		ID:       "12345678901234567890123456",
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}
	mock := mockUser{
		user:  &objCreated,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.GET("/users/:userId", uc.GetUser)
	id := "12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/"+id, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var objGet map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &objGet)
	assert.Nil(t, err)
	assert.Equal(t, objGet["id"], id)
	assert.Equal(t, objCreated.Name, objGet["name"])
	assert.Equal(t, objCreated.Username, objGet["username"])
	assert.Equal(t, "", objGet["password"])
	assert.Equal(t, objCreated.Email, objGet["email"])
}

func TestGetUser_NotFound(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.GET("/users/:userId", uc.GetUser)

	var uri = "/users/12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

}

func TestGetUser_ValidationError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.GET("/users/:userId", uc.GetUser)

	var uri = "/users/12345678"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"UserID\",\"message\":\"Field validation for 'UserID' failed on the 'len' tag\"}]", w.Body.String())
}

func TestGetUser_GenericError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   errors.New("generic error"),
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.GET("/users/:userId", uc.GetUser)

	var uri = "/users/12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestListUser(t *testing.T) {
	id := "12345678901234567890123456"
	objCreated := model.User{
		ID:       id,
		Username: "test",
		Password: "testsenha",
		Name:     "test nome",
		Email:    "test@mail.com",
	}
	mock := mockUser{
		user:  &objCreated,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()
	router.GET("/users", uc.ListUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var objSummary model.UserList // map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &objSummary)

	objList := objSummary.Data                    // ["data"].([]model.User)
	assert.Equal(t, int64(1), objSummary.Records) // ["records"].(int))

	for _, objGet := range objList {
		assert.Nil(t, err)
		assert.Equal(t, objGet.ID, id)
		assert.Equal(t, objCreated.Name, objGet.Name)
		assert.Equal(t, objCreated.Username, objGet.Username)
		assert.Equal(t, "", objGet.Password)
		assert.Equal(t, objCreated.Email, objGet.Email)
	}
}

func TestListUser_ValidationError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()
	router.GET("/users", uc.ListUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users?limit=abc", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Query\",\"message\":\"\\\"abc\\\" invalid syntax\"}]", w.Body.String())
}

func TestListUser_GenericError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   errors.New("generic error"),
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()
	router.GET("/users", uc.ListUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestUpdateUser(t *testing.T) {
	mock := mockUser{
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
	uc := UserController{mock}

	router := gin.Default()

	router.PUT("/users/:userId", uc.UpdateUser)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/12345678901234567890123456", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var objGet map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &objGet)
	assert.Nil(t, err)
	assert.NotNil(t, objGet["id"])
	assert.Len(t, objGet["id"], 26)
	assert.NotNil(t, objGet["name"])
	assert.NotNil(t, objGet["username"])
	assert.Equal(t, "", objGet["password"])
	assert.NotNil(t, objGet["email"])
}

func TestUpdateUser_ValidationError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.PUT("/users/:userId", uc.UpdateUser)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/12345678901234", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"UserID\",\"message\":\"Field validation for 'UserID' failed on the 'len' tag\"}]", w.Body.String())
}

func TestUpdateUser_BodyValidationError(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   nil,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.PUT("/users/:userId", uc.UpdateUser)

	b := bytes.NewReader([]byte(`{
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/12345678901234567890123456", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Username\",\"message\":\"Field validation for 'Username' failed on the 'required' tag\"}]", w.Body.String())
}

func TestUpdateUser_DuplicatedUser(t *testing.T) {
	mock := mockUser{
		user:  nil,
		err:   apierrors.DuplicatedUserError,
		count: 1,
	}
	uc := UserController{mock}

	router := gin.Default()

	router.PUT("/users/:userId", uc.UpdateUser)

	b := bytes.NewReader([]byte(`{
		"username": "test",
		"password":"testsenha",
		"name":"test nome",
		"email":"test@mail.com"
	}`))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/12345678901234567890123456", b)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "[{\"message\":\"Usuário já existe\"}]", w.Body.String())
}
