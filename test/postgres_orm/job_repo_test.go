package postgresorm_test

import (
	"testing"

	job_domain "github.com/rattapon001/porter-management/internal/job/domain"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	"github.com/stretchr/testify/assert"
)

func TestPGSave(t *testing.T) {
	assert := assert.New(t)
	db, err := ConnectionTestDB()
	if err != nil {
		t.Skip("TODO")
	}

	db.AutoMigrate(&job_domain.Job{})

	location := job_domain.Location{
		From: "Bangkok",
		To:   "Chiang Mai",
	}

	Patient := job_domain.Patient{
		Name: "Rattapon Prasongpongchai",
		HN:   "123456789",
	}
	job, err := job_domain.NewJob(location, Patient, nil)

	assert.Nil(err)
	jobRepository := job_postgres.NewPostgresOrmRepository(db)
	err = jobRepository.Save(job)
	assert.Nil(err)

	// defer db.Migrator().DropTable(&job_domain.Job{})
}
