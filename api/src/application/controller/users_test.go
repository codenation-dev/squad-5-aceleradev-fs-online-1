package controller

import (
	"app/resources/repository"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func initDB() *repository.UserRepository {
	db, err := xorm.NewEngine("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	repository.RunMigrations(db)

	return &repository.UserRepository{db}
}

func TestCreateUser(t *testing.T) {
	r := initDB()
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
}

func TestCreateUser_ValidationErrorUsername(t *testing.T) {
	r := initDB()
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

// func TestGetUser(t *testing.T) {
// 	type args struct {
// 		c *gin.Context
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			GetUser(tt.args.c)
// 		})
// 	}
// }
