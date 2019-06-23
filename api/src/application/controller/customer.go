package controller

import (
	"app/domain/builder"
	"app/domain/errors"
	"app/domain/service"
	"app/domain/validator"
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
		return
	}

	if f.Header.Get("Content-Type") != "text/csv" {
		errors.AbortWithError(c, &errors.ContentTypeInvalidError)
		return
	}

	fo, err := f.Open()

	if err != nil {
		errors.AbortWithError(c, &err)
		return
	}
	ci, err := cc.Customers.Parse(fo)

	if err != nil {
		errors.AbortWithError(c, &err)
		return
	}

	c.JSON(http.StatusOK, ci)

}

// CreateCustomer controller
func (cc CustomerController) CreateCustomer(c *gin.Context) {

	var customerCreate validator.CustomerCreation

	if err := c.ShouldBindJSON(&customerCreate); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	newCustomer := builder.CustomerCreationToCustomer(&customerCreate)

	customer, err := cc.Customers.CreateCustomer(newCustomer)

	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusCreated, customer)
	}
}

// UpdateCustomer - Atualiza um customer
func (cc CustomerController) UpdateCustomer(c *gin.Context) {
	var customerCreation validator.CustomerCreation
	if err := c.ShouldBindJSON(&customerCreation); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}
	var customerURI validator.CustomerURI
	if err := c.ShouldBindUri(&customerURI); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	customer, err := cc.Customers.UpdateCustomer(customerURI.CustomerID,
		builder.CustomerCreationToCustomer(&customerCreation))

	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusNoContent, customer)
	}
}

// ListCustomer Lista os customer
func (cc CustomerController) ListCustomer(c *gin.Context) {

	var q validator.CustomerListRequest

	if err := c.ShouldBindQuery(&q); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	customers, err := cc.Customers.ListCustomer(&q)

	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusOK, customers)
	}
}
