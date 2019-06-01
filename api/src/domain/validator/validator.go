package validator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// AbortWithValidation Abort with validation messages
func AbortWithValidation(c *gin.Context, err *error) {
	_, ok := (*err).(validator.ValidationErrors)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, *err)
		return
	}

	errors := make([]map[string]string, 0)
	for _, er := range (*err).(validator.ValidationErrors) {

		m := make(map[string]string)
		m["field"] = er.Field
		m["message"] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", er.Field, er.Tag)
		errors = append(errors, m)
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors)
}
