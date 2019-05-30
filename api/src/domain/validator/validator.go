package validator

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// AbortWithValidation Abort with validation messages
func AbortWithValidation(c *gin.Context, err *error) {
	switch t := (*err).(type) {
	case validator.ValidationErrors:
		validationErrors := (*err).(validator.ValidationErrors)
		errors := make([]map[string]string, 0)
		for _, er := range validationErrors {
			errors = append(errors, map[string]string{
				"field":   er.Field,
				"message": fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", er.Field, er.Tag),
			})
		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors)
	case *strconv.NumError:
		numError := (*err).(*strconv.NumError)
		msg := strings.Join([]string{"\"", numError.Num, "\" ", numError.Err.Error()}, "")
		errors := []map[string]string{map[string]string{"field": "Query", "message": msg}}
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors)
	default:
		log.Printf("Generic Validation error: %T\n", t)
		c.AbortWithError(http.StatusInternalServerError, *err)
	}
}
