package job_test

import (
	"context"
	"testing"

	"github.com/rattapon001/porter-management/internal/infra/uow"
	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
	"github.com/rattapon001/porter-management/pkg"
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

func (m *MockRepository) Update(job *domain.Job) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockRepository) FindById(id domain.JobId) (*domain.Job, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Job), args.Error(1)
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Publish(event []pkg.Event, uow uow.UnitOfWorkStore) error {
	args := m.Called(event)
	return args.Error(0)
}

func TestCreateNewJobUseCase(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Job")).Return(nil)
	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]pkg.Event")).Return(nil)
	JobUseCase := app.JobUseCaseImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}
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

	// publisher := mockPublisher
	createdJob, err := JobUseCase.CreateNewJob(context.Background(), location, patient, equipments)
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.JobPendingStatus, createdJob.Status, "created job status should be pending")
	assert.Equal(1, len(createdJob.Aggregate.Events), "created job should have 1 event")
	assert.Equal(2, len(createdJob.Equipments), "created job should have 2 equipments")
}
