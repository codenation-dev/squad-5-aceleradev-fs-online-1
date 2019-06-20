package email

import (
	"app/application/config"
)

var (
	Host     = config.GetEnv("EMAIL_HOST", "smtp.gmail.com")
	Port     = config.GetEnv("EMAIL_PORT", "465")
	Email    = config.GetEnv("EMAIL_USER", "banco.uati@gmail.com")
	Password = config.GetEnv("EMAIL_PASS", "bancouat123")
)
