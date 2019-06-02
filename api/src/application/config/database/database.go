package database

import "app/application/config"

var (
	// Host endereço de conexão ao banco de dados
	Host = config.GetEnv("DB_HOST", "db")
	// Port porta de conexão ao banco de dados
	Port = config.GetEnv("DB_PORT", "5432")
	// User usuário de conexão ao banco de dados
	User = config.GetEnv("DB_USER", "bancouat")
	// Password senha de conexão ao banco de dados
	Password = config.GetEnv("DB_PASSWORD", "bancouat")
	// DBname nome do banco de dados
	DBname = config.GetEnv("DB_NAME", "bancouat")
)
