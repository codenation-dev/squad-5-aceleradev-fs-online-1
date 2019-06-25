package service

import (
	jwtConfig "app/application/config/jwt"
	"app/domain/model"
	"app/domain/validator"
	"app/resources/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const expiresTime = 3600000000000 // Uma hora de validade

// Logins interfece
type Logins interface {
	Authorization(q validator.Login) (*model.Token, error)
}

// LoginService struct
type LoginService struct {
	Repository repository.LoginDB
}

// Authorization valida usuário e senha e gera token JWT
func (ls LoginService) Authorization(q validator.Login) (*model.Token, error) {

	err := ls.Repository.Authorization(q)

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + expiresTime,
		Id:        q.Username,
	})

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))

	if err != nil {
		return nil, err
	}

	return &model.Token{
		Token: tokenString,
	}, nil

}
