package validator

import (
	"app/domain/model"
	"reflect"
	"time"

	"gopkg.in/go-playground/validator.v8"
)

// IDURI struct
type IDURI struct {
	ID string `uri:"id" binding:"required,len=26"`
}

// AlertListRequest struct
type AlertListRequest struct {
	Name       string    `form:"name" binding:"omitempty,max=100"`
	Email      string    `form:"email" binding:"omitempty,max=160"`
	Customer   string    `form:"customer" binding:"omitempty,max=100"`
	Type       string    `form:"type" binding:"omitempty,alerttype"`
	DateStart  time.Time `form:"date_start" binding:"omitempty" time_format:"2006-01-02"`
	DateFinish time.Time `form:"date_finish" binding:"omitempty" time_format:"2006-01-02"`
	Limit      int       `form:"limit" binding:"omitempty,max=50"`
	Offset     int       `form:"offset" binding:"omitempty"`
}

// AlertTypeValidator valida o tipo AlertType
func AlertTypeValidator(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if t, ok := field.Interface().(string); ok {

		_, err := model.AlertTypeFromString(t)
		if err != nil {
			return false
		}
	}
	return true
}
