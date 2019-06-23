package jwt

import (
	"app/application/config"
)

var (
	// Secret palavra chave do JWT
	Secret = config.GetEnv("JWT_SECRET", "secret")
)
