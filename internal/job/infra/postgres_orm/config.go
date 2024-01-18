package postgresorm

type PostgresOrmConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func GetPostgresOrmConfig() *PostgresOrmConfig {
	return &PostgresOrmConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "123456",
		Database: "porter_management_db",
		SSLMode:  "disable",
	}
}
