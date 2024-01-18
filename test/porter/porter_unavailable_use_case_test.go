package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPorterUnavalable(t *testing.T) {
	assert := assert.New(t)
	porter, err := domain.NewPorter("porter1", "code-001", "token")
	assert.NoError(err, "should not return an error")

	mockRepo := new(MockRepository)
	mockRepo.On("FindByCode", mock.AnythingOfType("domain.PorterCode")).Return(porter)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Porter")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

	PorterUseCase := app.PorterUseCaseImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}
	porter, err = PorterUseCase.PorterUnavailable(porter.Code)
	assert.NoError(err, "should not return an error")
	assert.Equal(porter.Name, "porter1", "should return porter1")
	assert.Equal(porter.Status, domain.PorterStatusUnavailable, "should return available")
}
