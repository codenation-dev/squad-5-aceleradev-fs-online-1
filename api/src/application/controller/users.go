package controller

import (
	"app/domain/errors"
	"app/domain/validator"
	"app/resources/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser - Cadastrar um novo usuário
func CreateUser(r *repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var userCreation validator.UserCreation
		if err := c.ShouldBindJSON(&userCreation); err != nil {
			validator.AbortWithValidation(c, &err)
			return
		}

		user, err := r.CreateUser(&userCreation)

		if err != nil {
			errors.AbortWithError(c, &err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// GetUser - Consultar um usuário
func GetUser(r *repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var userURI validator.UserURI
		if err := c.ShouldBindUri(&userURI); err != nil {
			validator.AbortWithValidation(c, &err)
			return
		}

		user, err := r.GetUser(userURI.UserID)

		switch {
		case err != nil:
			errors.AbortWithError(c, &err)
		case user == nil:
			c.AbortWithStatus(http.StatusNotFound)
		default:
			c.JSON(http.StatusOK, user)
		}
	}
}
