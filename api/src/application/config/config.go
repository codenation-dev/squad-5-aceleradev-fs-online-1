package config

import "os"

// GetEnv recupera varivel de ambiente
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
