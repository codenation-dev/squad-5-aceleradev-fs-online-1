package builder

import (
	"app/domain/model"
	"app/domain/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreationToUser(t *testing.T) {
	args := validator.UserCreation{"user", "pass", "name", "test@mail.com"}
	want := model.User{
		Username: "user",
		Password: "pass",
		Name:     "name",
		Email:    "test@mail.com",
	}
	got := UserCreationToUser(&args)

	want.ID = got.ID

	assert.Equal(t, &want, got)

}
