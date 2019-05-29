package errors

import (
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_genericAPIError_Status(t *testing.T) {
	type fields struct {
		message string
		status  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "OK",
			fields: fields{
				message: "OK",
				status:  200,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := genericAPIError{
				message: tt.fields.message,
				status:  tt.fields.status,
			}
			if got := err.Status(); got != tt.want {
				t.Errorf("genericAPIError.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genericAPIError_Error(t *testing.T) {
	type fields struct {
		message string
		status  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				message: "OK",
				status:  200,
			},
			want: "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := genericAPIError{
				message: tt.fields.message,
				status:  tt.fields.status,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("genericAPIError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIValidationError(t *testing.T) {
	type args struct {
		status int
		msg    string
	}
	tests := []struct {
		name string
		args args
		want APIValidationError
	}{
		{
			name: "OK",
			args: args{
				status: 200,
				msg:    "OK",
			},
			want: genericAPIError{"OK", 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPIValidationError(tt.args.status, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbortWithError(t *testing.T) {
	type args struct {
	}
	err := errors.New("test")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	AbortWithError(c, &err)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestAbortWithError_APIValidation(t *testing.T) {
	type args struct {
	}
	var err error = NewAPIValidationError(400, "erro de validação")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	AbortWithError(c, &err)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "[{\"message\":\"erro de validação\"}]", w.Body.String())
}
