package postgresorm

import (
	"os"
	"sync"
)

type PostgresOrmConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

var pgConfig *PostgresOrmConfig
var once sync.Once

func GetPostgresOrmConfig() *PostgresOrmConfig {
	once.Do(func() {
		pgConfig = &PostgresOrmConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		}
	})
	return pgConfig
}
