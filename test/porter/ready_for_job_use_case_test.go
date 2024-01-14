package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestJobAvailable(t *testing.T) {
	assert := assert.New(t)
	porter, err := domain.CreatedNewPorter("porter1", "code-001")
	assert.NoError(err, "should not return an error")

	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.AnythingOfType("*domain.Porter")).Return(nil)

	mockPublisher := new(MockEventHandler)
	mockPublisher.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

	porterService := app.PorterServiceImpl{
		Repo:      mockRepo,
		Publisher: mockPublisher,
	}
	err = porterService.ReadyForJob(porter, "code-001")
	assert.NoError(err, "should not return an error")
}
