package errors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIValidationError interface
type APIValidationError interface {
	Status() int
	Error() string
}

type genericAPIError struct {
	message string
	status  int
}

func (err genericAPIError) Status() int {
	return err.status
}

func (err genericAPIError) Error() string {
	return err.message
}

// NewAPIValidationError cria um novo tipo de erro
func NewAPIValidationError(status int, msg string) APIValidationError {
	return genericAPIError{
		message: msg,
		status:  status,
	}
}

// AbortWithError Abort with error messages
func AbortWithError(c *gin.Context, err *error) {
	errorAPI, ok := (*err).(APIValidationError)
	if ok && err != nil {
		log.Println("APIValidationError", err)

		errors := make([]map[string]string, 1, 1)
		m := make(map[string]string)
		m["message"] = errorAPI.Error()
		errors[0] = m
		c.AbortWithStatusJSON(errorAPI.Status(), errors)
	} else {
		log.Println("Error", err)
		c.AbortWithError(http.StatusInternalServerError, *err)
	}
}
