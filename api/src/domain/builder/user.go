package builder

import (
	"app/domain/model"
	"app/domain/validator"
)

// UserCreationToUser converte UserCreation para User
func UserCreationToUser(userCreation *validator.UserCreation) *model.User {
	return &model.User{
		ID:       NewULID(),
		Username: userCreation.Username,
		Password: userCreation.Password,
		Name:     userCreation.Name,
		Email:    userCreation.Email,
	}
}

// GetRecipients retorna a lista de e-mails
func GetRecipients(users []model.User) []string {
	s := len(users)
	emails := make([]string, s, s)
	for i, u := range users {
		emails[i] = u.Email
	}
	return emails
}
