package controller

import (
	"app/domain/errors"
	"app/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomerController struct
type CustomerController struct {
	Customers service.Customer
}

// UploadCustomer faz o upload da lista de customers e grava no banco
func (cc CustomerController) UploadCustomer(c *gin.Context) {

	f, err := c.FormFile("file")

	if err != nil {
		errors.AbortWithError(c, &err)
	}

	if f.Header.Get("Content-Type") != "text/csv" {
		errors.AbortWithError(c, &errors.ContentTypeInvalidError)
	}

	fo, err := f.Open()

	if err != nil {
		errors.AbortWithError(c, &err)
	}
	cl, err := cc.Customers.Parse(fo)

	if err != nil {
		errors.AbortWithError(c, &err)
	}

	c.JSON(http.StatusOK, cl)

}
