package config

import (
	"log"
	"os"
)

// GetEnv recupera varivel de ambiente
func GetEnv(key ...string) string {
	if value, exists := os.LookupEnv(key[0]); exists {
		return value
	}

	if len(key) <= 1 {
		log.Printf("Variavel de ambiente '%s' não foi definida e não tem valor padrão\n", key[0])
		panic(1)
	}

	return key[1]
}
