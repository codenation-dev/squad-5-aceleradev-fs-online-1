package builder

import (
	"app/domain/model"
	"app/domain/validator"
)

// UserCreationToUser converte UserCreation para User
func UserCreationToUser(userCreation *validator.UserCreation) *model.User {
	return &model.User{
		ID:       newULID(),
		Username: userCreation.Username,
		Password: userCreation.Password,
		Name:     userCreation.Name,
		Email:    userCreation.Email,
	}
}
