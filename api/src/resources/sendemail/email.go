package sendemail

import (
	"app/application/config/email"
	"crypto/tls"
	"net/smtp"
)

// ClientEmail strct
type ClientEmail struct {
}

//NewClientEmail atribui um *smtp.Client
func (c ClientEmail) NewClientEmail() (*smtp.Client, error) {

	auth := smtp.PlainAuth("", email.Email, email.Password, email.Host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         email.Host,
	}

	conn, err := tls.Dial("tcp", serverName(), tlsconfig)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, email.Host)
	if err != nil {
		return nil, err
	}

	if err = client.Auth(auth); err != nil {
		return nil, err
	}

	// step 2: add all from and to
	if err = client.Mail(email.Email); err != nil {
		return nil, err
	}

	return client, nil
}

func serverName() string {
	return email.Host + ":" + email.Port
}
