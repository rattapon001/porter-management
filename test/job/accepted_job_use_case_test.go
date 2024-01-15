package job_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/rattapon001/porter-management/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAcceptedJob(t *testing.T) {
	assert := assert.New(t)

	mockRepo := new(MockRepository)
	mockRepo.On("Update", mock.AnythingOfType("*domain.Job")).Return(nil)
	mockRepo.On("FindById", mock.AnythingOfType("domain.JobId")).Return(&domain.Job{
		Status: domain.JobStatusPending,
		Aggregate: domain.Aggregate{
			Events: []pkg.Event{
				jobCreatedEvent,
			},
		},
	}, nil)
	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]pkg.Event")).Return(nil)
	jobService := app.JobServiceImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}
	acceptedJob, err := jobService.AcceptedJob("1", porter)
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.JobStatusAccepted, acceptedJob.Status, "accepted job status should be accepted")
	assert.Equal(2, len(acceptedJob.Aggregate.Events), "accepted job should have 2 events")
}
