package job_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(job *domain.Job) error {
	args := m.Called(job)
	return args.Error(0)
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Publish(event []domain.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func TestCreatedNewJob(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Job")).Return(nil)
	jobService := app.JobServiceImpl{
		Repo: mockRepo,
	}
	location := domain.Location{
		From: "A",
		To:   "B",
	}
	patient := domain.Patient{
		Name: "John Smith",
		HN:   "HN123",
	}
	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]domain.Event")).Return(nil)
	publisher := mockPublisher
	createdJob, err := jobService.CreatedNewJob(location, patient, publisher)
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.JobStatusPending, createdJob.Status, "created job status should be pending")
	assert.Equal(1, len(createdJob.Aggregate.Events), "created job should have 1 event")
}
