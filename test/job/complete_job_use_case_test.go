package job_test

import (
	"context"
	"log"
	"testing"

	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/rattapon001/porter-management/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompleteJobUseCase(t *testing.T) {
	assert := assert.New(t)

	mockRepo := new(MockRepository)
	mockRepo.On("FindById", mock.AnythingOfType("domain.JobId")).Return(&domain.Job{
		Status: domain.JobWorkingStatus,
		Aggregate: domain.Aggregate{
			Events: []pkg.Event{
				jobCreatedEvent, jobAcceptedEvent, jobStartedEvent,
			},
		},
	}, nil)

	mockRepo.On("Save", mock.AnythingOfType("*domain.Job")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]pkg.Event")).Return(nil)
	JobUseCase := app.JobUseCaseImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}

	startedJob, err := JobUseCase.CompleteJob(context.Background(), "1")
	log.Println("TestCompletedJob startedJob", startedJob)
	log.Println("TestCompletedJob err", err)
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.JobCompletedStatus, startedJob.Status, "completed job status should be completed")
	assert.Equal(4, len(startedJob.Aggregate.Events), "completed job should have 4 events")

}
