package controller

import (
	"app/domain/validator"
	"app/resources/repository"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func initDB(runMigrations bool) *repository.UserRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if runMigrations {
		repository.RunMigrations(db)
	}

	return &repository.UserRepository{DB: db}
}

func TestCreateUser(t *testing.T) {
	r := initDB(true)
	defer r.DB.Close()
	router := gin.Default()

	router.POST("/users", CreateUser(r))

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
	r := initDB(true)
	defer r.DB.Close()
	router := gin.Default()

	router.POST("/users", CreateUser(r))

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
	r := initDB(false)
	defer r.DB.Close()
	router := gin.Default()

	router.POST("/users", CreateUser(r))

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
	r := initDB(true)
	defer r.DB.Close()
	router := gin.Default()

	_, err := r.CreateUser(&validator.UserCreation{
		Username: "test",
		Password: "test",
		Name:     "test nome",
		Email:    "test@mail.com",
	})
	assert.Nil(t, err)

	router.POST("/users", CreateUser(r))

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
	r := initDB(true)
	defer r.DB.Close()
	router := gin.Default()

	router.POST("/users", CreateUser(r))
	router.GET("/users/:userId", GetUser(r))

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

	var objCreated map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &objCreated)
	assert.Nil(t, err)

	id, ok := objCreated["id"]
	assert.True(t, ok)
	assert.Len(t, id, 26)

	var uri = "/users/" + id

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var objGet map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &objGet)
	assert.Nil(t, err)
	assert.Equal(t, objGet["id"], id)
	assert.Equal(t, objCreated["name"], objGet["name"])
	assert.Equal(t, objCreated["username"], objGet["username"])
	assert.Equal(t, "", objGet["password"])
	assert.Equal(t, objCreated["email"], objGet["email"])
}

func TestGetUser_NotFound(t *testing.T) {
	r := initDB(true)
	defer r.DB.Close()
	router := gin.Default()

	router.GET("/users/:userId", GetUser(r))

	var uri = "/users/12345678901234567890123456"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

}

func TestGetUser_ValidationError(t *testing.T) {
	r := initDB(true)
	defer r.DB.Close()
	router := gin.Default()

	router.GET("/users/:userId", GetUser(r))

	var uri = "/users/12345678"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"UserID\",\"message\":\"Field validation for 'UserID' failed on the 'len' tag\"}]", w.Body.String())
}
