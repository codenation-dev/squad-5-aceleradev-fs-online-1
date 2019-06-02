package controller

import (
	"app/domain/service"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Customer service.Customer
}

func (uc CustomerController) UploadCustomer(c *gin.Context) {


}
