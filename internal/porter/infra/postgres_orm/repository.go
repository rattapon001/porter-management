package postgresorm

import "gorm.io/gorm"

type PostgresOrmRepository struct {
	db *gorm.DB
}

func NewPostgresOrmRepository(db *gorm.DB) *PostgresOrmRepository {
	return &PostgresOrmRepository{
		db: db,
	}
}
