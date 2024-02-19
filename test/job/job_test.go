package job_test

import (
	"testing"
	"time"

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
	EventName: domain.JobCreatedEvent,
	Payload: map[string]interface{}{
		"job_id":   uuid.New().String(),
		"version":  1,
		"status":   domain.JobPendingStatus,
		"location": location,
		"patient":  patient,
	},
}

var jobAcceptedEvent = pkg.Event{
	EventName: domain.JobCreatedEvent,
	Payload: map[string]interface{}{
		"job_id":   uuid.New().String(),
		"version":  2,
		"status":   domain.JobAcceptedStatus,
		"location": location,
		"patient":  patient,
		"porter":   porter,
	},
}

var jobStartedEvent = pkg.Event{
	EventName: domain.JobCreatedEvent,
	Payload: map[string]interface{}{
		"job_id":   uuid.New().String(),
		"version":  3,
		"status":   domain.JobWorkingStatus,
		"location": location,
		"patient":  patient,
		"porter":   porter,
		"check_in": time.Now(),
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

	equipments := []domain.Equipment{
		{
			EquipmentId: "1",
			Amount:      1,
		},
		{
			EquipmentId: "2",
			Amount:      1,
		},
	}

	createdJob, err := domain.NewJob(location, patient, equipments)
	assert.Nil(err, "error should be nil")
	assert.Equal(domain.JobPendingStatus, createdJob.Status, "created job status should be pending")
	assert.Equal(1, len(createdJob.Aggregate.Events), "created job should have 1 event")
}
