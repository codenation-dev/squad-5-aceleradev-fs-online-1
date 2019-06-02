package controller

import (
	"app/domain/service"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Customer service.Customer
}

func (uc CustomerController) UpdateCustomer(c *gin.Context) {

	var cs service.CustomerService

	f, err := c.FormFile("file")

	if err != nil {
		c.Status(500)
	}
	fo, err := f.Open()
	cs.Parse(fo)
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": f.Header.Get("Content-Type"),
	})

}
