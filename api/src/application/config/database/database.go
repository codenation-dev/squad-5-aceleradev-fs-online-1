package database

import "app/application/config"

var (
	Host     = config.GetEnv("DB_HOST", "db")
	Port     = config.GetEnv("DB_PORT", "5432")
	User     = config.GetEnv("DB_USER", "bancouat")
	Password = config.GetEnv("DB_PASSWORD", "bancouat")
	DBname   = config.GetEnv("DB_NAME", "bancouat")
)
