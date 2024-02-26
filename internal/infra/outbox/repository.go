package outbox

import "gorm.io/gorm"

type OutboxRepository interface {
	Save(event OutboxEvent) error
}

type PostgresOutboxRepository struct {
	db *gorm.DB
}

func NewPostgresOutboxRepository(db *gorm.DB) *PostgresOutboxRepository {
	return &PostgresOutboxRepository{db: db}
}

func (por *PostgresOutboxRepository) Save(event OutboxEvent) error {
	return por.db.Create(&event).Error
}
