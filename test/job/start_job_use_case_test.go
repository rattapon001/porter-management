package job_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/rattapon001/porter-management/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStartedJob(t *testing.T) {
	assert := assert.New(t)

	mockRepo := new(MockRepository)
	mockRepo.On("FindById", mock.AnythingOfType("domain.JobId")).Return(&domain.Job{
		Status: domain.JobStatusAccepted,
		Aggregate: domain.Aggregate{
			Events: []pkg.Event{
				jobCreatedEvent, jobAcceptedEvent,
			},
		},
	}, nil)
	mockRepo.On("Update", mock.AnythingOfType("*domain.Job")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]pkg.Event")).Return(nil)
	JobUseCase := app.JobUseCaseImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}

	startedJob, err := JobUseCase.StartJob("1")
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.JobStatusWorking, startedJob.Status, "started job status should be started")
	assert.Equal(3, len(startedJob.Aggregate.Events), "started job should have 3 events")

}
