package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPorterWorking(t *testing.T) {
	assert := assert.New(t)
	porter, err := domain.CreatedNewPorter("porter1", "code-001", "token")
	assert.NoError(err, "should not return an error")

	mockRepo := new(MockRepository)
	mockRepo.On("FindByID", mock.AnythingOfType("domain.PorterId")).Return(porter)
	mockRepo.On("Update", mock.AnythingOfType("*domain.Porter")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

	porterService := app.PorterServiceImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}
	porter, err = porterService.PorterWorking(porter.ID)
	assert.NoError(err, "should not return an error")
	assert.Equal(porter.Name, "porter1", "should return porter1")
	assert.Equal(porter.Status, domain.PorterStatusWorking, "should return working")
}
