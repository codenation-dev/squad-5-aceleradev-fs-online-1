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

func initDBUser(runMigrations bool) *UserRepository {
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
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	newUser, err := uRepo.GetUser("1111")
	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, "1111", newUser.ID)
	assert.Equal(t, "test", newUser.Username)
	assert.Equal(t, "test", newUser.Password)
	assert.Equal(t, "test", newUser.Name)
	assert.Equal(t, "test@mail.com", newUser.Email)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)
}

func TestUserRepository_GetUser_NotFound(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	newUser, err := uRepo.GetUser("1111")
	assert.Nil(t, err)
	assert.Nil(t, newUser)
}

func TestUserRepository_CreateUser(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	newUser := model.User{
		ID:       "1111",
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}

	err := uRepo.CreateUser(&newUser)
	assert.Nil(t, err)

	user := model.User{}
	ok, err := uRepo.DB.ID(newUser.ID).Get(&user)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, ok)
	newUser.CreatedAt = user.CreatedAt
	newUser.UpdatedAt = user.UpdatedAt
	assert.Equal(t, user, newUser)
}

func TestUserRepository_CreateUser_DuplicatedUser(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	newUser := model.User{
		ID:       "1111",
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}

	err := uRepo.CreateUser(&newUser)
	assert.Nil(t, err)

	user := model.User{}
	ok, err := uRepo.DB.ID(newUser.ID).Get(&user)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, ok)
	newUser.CreatedAt = user.CreatedAt
	newUser.UpdatedAt = user.UpdatedAt
	assert.Equal(t, user, newUser)

	err = uRepo.CreateUser(&newUser)
	assert.NotNil(t, err)
	assert.Equal(t, errors.DuplicatedUserError, err)
}

func TestUserRepository_CountUsers(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{}

	c, err := uRepo.CountUsers(&q)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), c)
}

func TestUserRepository_CountUsers_Name(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{Name: "test"}

	c, err := uRepo.CountUsers(&q)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)
}

func TestUserRepository_CountUsers_Email(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{Email: "test@mail.com"}

	c, err := uRepo.CountUsers(&q)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), c)
}

func TestUserRepository_ListUser(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{}

	list, err := uRepo.ListUser(&q)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(*list))
}

func TestUserRepository_ListUser_Name(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{Name: "test"}

	list, err := uRepo.ListUser(&q)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*list))
}

func TestUserRepository_ListUser_Email(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{Email: "test@mail.com"}

	list, err := uRepo.ListUser(&q)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*list))
}

func TestUserRepository_ListUser_Limit(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{Limit: 1}

	list, err := uRepo.ListUser(&q)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*list))

	newUser := (*list)[0]
	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, "1111", newUser.ID)
	assert.Equal(t, "test", newUser.Username)
	assert.Equal(t, "test", newUser.Password)
	assert.Equal(t, "test", newUser.Name)
	assert.Equal(t, "test@mail.com", newUser.Email)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)
}

func TestUserRepository_ListUser_Offset(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	q := validator.UserListRequest{Offset: 1}

	list, err := uRepo.ListUser(&q)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*list))

	newUser := (*list)[0]
	assert.Nil(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, "2222", newUser.ID)
	assert.Equal(t, "bla", newUser.Username)
	assert.Equal(t, "bla", newUser.Password)
	assert.Equal(t, "bla", newUser.Name)
	assert.Equal(t, "bla@mail.com", newUser.Email)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)

}

func TestUserRepository_UpdateUser(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	uc := model.User{
		ID:       "1111",
		Username: "test2",
		Password: "test",
		Name:     "test2",
		Email:    "test2@mail.com",
	}

	err := uRepo.UpdateUser(&uc)
	assert.Nil(t, err)

	newUser := model.User{}
	ok, err := uRepo.DB.ID(uc.ID).Get(&newUser)
	assert.Nil(t, err)
	assert.True(t, ok)

	assert.NotNil(t, newUser)
	assert.Equal(t, uc.Username, newUser.Username)
	assert.Equal(t, uc.Password, newUser.Password)
	assert.Equal(t, uc.Name, newUser.Name)
	assert.Equal(t, uc.Email, newUser.Email)
	assert.NotNil(t, newUser.ID)
	assert.NotNil(t, newUser.CreatedAt)
	assert.NotNil(t, newUser.UpdatedAt)
}

func TestUserRepository_UpdateUser_DuplicatedUser(t *testing.T) {
	uRepo := initDBUser(true)
	defer uRepo.DB.Close()

	mockUsers(uRepo.DB)

	uc := model.User{
		ID:       "1111",
		Username: "bla",
		Password: "test",
		Name:     "test2",
		Email:    "test2@mail.com",
	}

	err := uRepo.UpdateUser(&uc)
	assert.NotNil(t, err)
	assert.Equal(t, errors.DuplicatedUserError, err)
}

func mockUsers(DB *xorm.Engine) {
	u := model.User{
		ID:       "1111",
		Username: "test",
		Password: "test",
		Name:     "test",
		Email:    "test@mail.com",
	}
	DB.InsertOne(u)

	u = model.User{
		ID:       "2222",
		Username: "bla",
		Password: "bla",
		Name:     "bla",
		Email:    "bla@mail.com",
	}
	DB.InsertOne(u)
}
