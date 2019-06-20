package service

import (
	"app/domain/model"
	"app/resources/sendemail"
)

// Email interface
type Email interface {
	Send(e model.Email) error
}

// EmailService struct
type EmailService struct {
	Client sendemail.ClientEmail
}

// Send envia um email
func (e EmailService) Send(email model.Email) error {

	es := EmailService{}
	c, err := es.Client.NewClientEmail()
	defer c.Quit()
	if err != nil {
		return err
	}

	for _, k := range email.Recipients {
		c.Rcpt(k)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write([]byte("Subject:" + email.Subject + "\n\n" + email.Body))
	if err != nil {
		return err
	}

	return nil
}
