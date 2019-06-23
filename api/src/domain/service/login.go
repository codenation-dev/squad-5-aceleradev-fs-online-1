package service

import (
	jwtConfig "app/application/config/jwt"
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"

	"github.com/dgrijalva/jwt-go"
)

// Logins interfece
type Logins interface {
	Authorization(q validator.Login) (*model.Token, error)
}

// LoginService struct
type LoginService struct {
	Repository repository.LoginDB
}

func (ls LoginService) Authorization(q validator.Login) (*model.Token, error) {

	err := ls.Repository.Authorization(q)

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))

	if err != nil {
		return nil, err
	}

	return &model.Token{
		Token: tokenString,
	}, nil

}
