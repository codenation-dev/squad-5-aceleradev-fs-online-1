package controller

import (
	"app/domain/errors"
	"app/domain/service"
	"app/domain/validator"
	"app/resources/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct {
	Users service.Users
}

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

// ListUser - Listar os usuários
func (uc UserController) ListUser(c *gin.Context) {
	var q validator.UserListRequest
	if err := c.ShouldBindQuery(&q); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	users, err := uc.Users.ListUsers(&q)
	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// UpdateUser - Atualiza um usuário
func (uc UserController) UpdateUser(c *gin.Context) {
	var userCreation validator.UserCreation
	if err := c.ShouldBindJSON(&userCreation); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}
	var userURI validator.UserURI
	if err := c.ShouldBindUri(&userURI); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	user, err := uc.Users.UpdateUser(userURI.UserID, &userCreation)

	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
