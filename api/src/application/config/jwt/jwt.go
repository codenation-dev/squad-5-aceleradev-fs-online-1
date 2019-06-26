package jwt

import (
	"app/application/config"
)

var (
	// Secret palavra chave do JWT
	Secret = config.GetEnv("JWT_SECRET", "secret")
	// Disabled indica se a validação do token está ligada
	Disabled = config.GetEnv("JWT_DISABLED", "false")
)
