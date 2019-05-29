package validator

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v8"
)

func TestAbortWithValidation(t *testing.T) {
	type args struct {
	}
	err := errors.New("test")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	AbortWithValidation(c, &err)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestAbortWithValidation_ValidationErrors(t *testing.T) {
	type args struct {
	}
	var err error = validator.ValidationErrors{"field": &validator.FieldError{Field: "field", Tag: "len"}}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	AbortWithValidation(c, &err)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"field\",\"message\":\"Field validation for 'field' failed on the 'len' tag\"}]", w.Body.String())
}
