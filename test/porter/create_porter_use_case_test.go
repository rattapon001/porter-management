package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/rattapon001/porter-management/pkg"
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

func (m *MockRepository) Update(job *domain.Porter) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockRepository) FindAvailablePorter() *domain.Porter {
	args := m.Called()
	return args.Get(0).(*domain.Porter)
}

func (m *MockRepository) FindByID(id domain.PorterId) *domain.Porter {
	args := m.Called(id)
	return args.Get(0).(*domain.Porter)
}

func (m *MockRepository) FindByCode(code domain.PorterCode) *domain.Porter {
	args := m.Called(code)
	return args.Get(0).(*domain.Porter)
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Publish(event []pkg.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func TestCreatePorterUseCase(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Porter")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

	PorterUseCase := app.PorterUseCaseImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}

	token := "token"

	porter, err := PorterUseCase.CreateNewPorter("porter1", token)
	assert.NoError(err, "should not return an error")
	assert.Equal(domain.PorterStatusUnavailable, porter.Status, "created porter status should be available")
}
