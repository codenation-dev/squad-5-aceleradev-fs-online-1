package controller

import (
	"app/domain/errors"
	"app/domain/service"
	"app/domain/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginController struct
type LoginController struct {
	Login service.Logins
}

func (lc LoginController) Authorization(c *gin.Context) {

	var q validator.Login

	if err := c.ShouldBindJSON(&q); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	token, err := lc.Login.Authorization(q)

	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusOK, token)
	}
}
