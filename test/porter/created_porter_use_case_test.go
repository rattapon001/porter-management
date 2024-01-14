package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/app"
	event "github.com/rattapon001/porter-management/internal/porter/app/event_handler"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(job *domain.Porter) error {
	args := m.Called(job)
	return args.Error(0)
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Publish(event []event.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func TestCreatedPorterUseCase(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Porter")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

	porterService := app.PorterServiceImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}

	porter, err := porterService.CreatedNewPorter("porter1")
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.PorterStatusUnavailable, porter.Status, "created porter status should be available")
}
