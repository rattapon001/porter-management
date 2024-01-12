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

	createdJob, err := domain.CreateNewJob(location, patient)
	assert.Nil(err, "error should be nil")
	assert.Equal(domain.JobStatusPending, createdJob.Status, "created job status should be pending")
	assert.Equal(1, len(createdJob.Aggregate.Events), "created job should have 1 event")
}
