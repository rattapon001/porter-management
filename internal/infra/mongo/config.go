package mongo_db

import (
	"os"
	"sync"
)

type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var mongoConfig *MongoConfig
var once sync.Once

func GetMongoConfig() *MongoConfig {
	once.Do(func() {
		mongoConfig = &MongoConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		}
	})
	return mongoConfig
}
