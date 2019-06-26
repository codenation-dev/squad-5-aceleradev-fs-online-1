package email

import (
	"app/application/config"
)

var (
	// Host endereço de conexão ao servidor de e-mail
	Host = config.GetEnv("EMAIL_HOST", "smtp.gmail.com")
	// Port porta de conexão ao servidor de e-mail
	Port = config.GetEnv("EMAIL_PORT", "465")
	// Email conexão ao servidor de e-mail
	Email = config.GetEnv("EMAIL_USER", "banco.uati@gmail.com")
	// Password senha conexão ao servidor de e-mail
	Password = config.GetEnv("EMAIL_PASS", "bancouat123")
	// Disabled envio de e-mail desabilitado
	Disabled = config.GetEnv("EMAIL_DISABLED", "false")
)
