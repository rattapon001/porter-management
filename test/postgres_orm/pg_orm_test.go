package postgresorm_test

import (
	"github.com/joho/godotenv"
	postgresorm "github.com/rattapon001/porter-management/internal/infra/postgres_orm"
	job_domain "github.com/rattapon001/porter-management/internal/job/domain"
	porter_domain "github.com/rattapon001/porter-management/internal/porter/domain"
	"gorm.io/gorm"
)

func ConnectionTestDB() (*gorm.DB, error) {
	err := godotenv.Load("../../configs/.env.test")
	if err != nil {
		return nil, err
	}

	config := postgresorm.GetPostgresOrmConfig()
	db, err := postgresorm.NewPostgresOrmDbConnection(config)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&job_domain.Job{}, &porter_domain.Porter{})
	return db, nil
}
