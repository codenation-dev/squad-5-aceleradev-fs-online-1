package sendemail

import (
	"app/application/config/email"
	"github.com/stretchr/testify/assert"
	"net/smtp"
	"reflect"
	"testing"
)

func TestClientEmail_NewClientEmail(t *testing.T) {
	tests := []struct {
		name    string
		c       ClientEmail
		want    *smtp.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ClientEmail{}
			got, err := c.NewClientEmail()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientEmail.NewClientEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientEmail.NewClientEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientEmail_NewClientEmail_ErroAuth(t *testing.T) {

	email.Password = "1234"

	c := ClientEmail{}

	_, err := c.NewClientEmail()

	assert.NotNil(t, err)
}

func TestClientEmail_NewClientEmail_ErroTCP(t *testing.T) {

	email.Host = "1234"

	c := ClientEmail{}

	_, err := c.NewClientEmail()

	assert.NotNil(t, err)
}

func TestClientEmail_NewClientEmail_ErroSmtpNewClient(t *testing.T) {

	email.Host = "1234"

	c := ClientEmail{}

	_, err := c.NewClientEmail()

	assert.NotNil(t, err)
}

func TestClientEmail_NewClientEmail_ErroMail(t *testing.T) {

	email.Email = "1234"

	c := ClientEmail{}

	_, err := c.NewClientEmail()

	assert.NotNil(t, err)
}
