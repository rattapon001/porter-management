package mongo_db

type MongoConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetMongoConfig() *MongoConfig {
	return &MongoConfig{
		Host:     "localhost",
		Port:     "27017",
		User:     "root",
		Password: "123456",
		Database: "porter_management_db",
	}
}
