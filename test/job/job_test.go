package job_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewJob(t *testing.T) {
	assert := assert.New(t)

	location := domain.Location{
		From: "A",
		To:   "B",
	}
	patient := domain.Patient{
		Name: "John Smith",
		HN:   "HN123",
	}
	porter := domain.Porter{
		Code: "P123",
		Name: "Porter",
	}

	createdJob := domain.CreateNewJob(location, patient, porter)
	assert.Equal(domain.JobStatusPending, createdJob.Status, "created job status should be pending")
	assert.Equal(1, len(createdJob.Aggregate.Events), "created job should have 1 event")
}
