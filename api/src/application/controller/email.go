package controller

import (
	"app/domain/model"
	"app/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailController struct {
	EmailService service.EmailService
}

// Send envia o email
func (ec *EmailController) Send(c *gin.Context) {

	email := model.Email{
		Recipients: []string{"viniciosampaio@hotmail.com", "viniciosampaio@gmail.com"},
		Subject:    "Subject:Olá!\n\n",
	}

	err := ec.EmailService.Send(email)

	c.JSON(http.StatusOK, err)

}
