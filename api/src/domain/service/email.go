package service

import (
	config "app/application/config/email"
	"app/domain/model"
	"app/resources/sendemail"
	"log"
	"strconv"
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
	if ok, err := strconv.ParseBool(config.Disabled); ok && err == nil {
		log.Printf("Email Send %#v", email)
		return nil
	}

	es := EmailService{}
	c, err := es.Client.NewClientEmail()
	if err != nil {
		return err
	}
	defer c.Quit()

	for _, k := range email.Recipients {
		err := c.Rcpt(k)
		if err != nil {
			return err
		}
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
