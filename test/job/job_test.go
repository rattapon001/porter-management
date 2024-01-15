package job_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/rattapon001/porter-management/pkg"
	"github.com/stretchr/testify/assert"
)

var location = domain.Location{
	From: "A",
	To:   "B",
}

var patient = domain.Patient{
	Name: "John Smith",
	HN:   "HN123",
}

var porter = domain.Porter{
	Name: "Porter",
	Code: "P001",
}

var jobCreatedEvent = pkg.Event{
	EventName: domain.JobEventCreated,
	Payload: map[string]interface{}{
		"job_id":   uuid.New().String(),
		"version":  1,
		"status":   domain.JobStatusPending,
		"location": location,
		"patient":  patient,
	},
}

var jobAcceptedEvent = pkg.Event{
	EventName: domain.JobEventCreated,
	Payload: map[string]interface{}{
		"job_id":   uuid.New().String(),
		"version":  2,
		"status":   domain.JobStatusAccepted,
		"location": location,
		"patient":  patient,
		"porter":   porter,
	},
}

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

	createdJob, err := domain.CreatedNewJob(location, patient)
	assert.Nil(err, "error should be nil")
	assert.Equal(domain.JobStatusPending, createdJob.Status, "created job status should be pending")
	assert.Equal(1, len(createdJob.Aggregate.Events), "created job should have 1 event")
}
