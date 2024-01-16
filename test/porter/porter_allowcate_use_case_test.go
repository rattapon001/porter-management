package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/rattapon001/porter-management/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNotification struct {
	mock.Mock
}

func (m *MockNotification) Notify(token string, payload pkg.NotificationPayload) error {
	args := m.Called(token, payload)
	return args.Error(0)
}

func TestPorterAllowcate(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockRepository)

	porter, _ := domain.NewPorter("porter1", "code-001", "token")
	mockRepo.On("FindAvailablePorter").Return(porter)
	mockNoti := new(MockNotification)
	mockNoti.On("Notify", mock.AnythingOfType("string"), mock.AnythingOfType("pkg.NotificationPayload")).Return(nil)
	porterService := app.PorterServiceImpl{
		Repo: mockRepo,
		Noti: mockNoti,
	}
	location := domain.Location{
		From: "A",
		To:   "B",
	}
	patient := domain.Patient{
		Name: "John Smith",
		HN:   "HN123",
	}
	job := domain.Job{
		Location: location,
		Patient:  patient,
	}
	porter, err := porterService.PorterAllowcate(job)
	assert.NoError(err, "should not return an error")
	assert.Equal(porter.Name, "porter1", "should return porter1")
}
