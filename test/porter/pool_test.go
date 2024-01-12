package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreatedNewPool(t *testing.T) {
	assert := assert.New(t)

	pool := domain.CreateNewPool()
	assert.Equal(0, len(pool.Porters), "pool should be empty")
}
func TestPorterRegister(t *testing.T) {
	assert := assert.New(t)

	pool := domain.CreateNewPool()
	porter, err := domain.CreatedNewPorter("John Smith", "porter-001")
	assert.Nil(err, "error should be nil")
	pool.PorterRegister(porter, "token-001")
	assert.Equal(1, len(pool.Porters), "pool should have 1 porter")
	assert.Equal(domain.PorterStatusAvailable, pool.Porters[0].Status, "porter status should be available")
	assert.Equal("token-001", pool.Porters[0].Token, "porter token should be token-001")
}

func TestFindAvailablePorter(t *testing.T) {
	assert := assert.New(t)

	pool := domain.CreateNewPool()
	porter, err := domain.CreatedNewPorter("John Smith", "porter-001")
	assert.Nil(err, "error should be nil")
	pool.PorterRegister(porter, "token-001")
	availablePorter := pool.FindAvailablePorter()
	assert.Equal(porter.ID, availablePorter.ID, "porter should be available")
	assert.Equal(domain.PorterStatusAvailable, availablePorter.Status, "porter status should be available")
}

func TestPorterAcceptJob(t *testing.T) {
	assert := assert.New(t)

	pool := domain.CreateNewPool()
	porter, err := domain.CreatedNewPorter("John Smith", "porter-001")
	assert.Nil(err, "error should be nil")
	pool.PorterRegister(porter, "token-001")
	pool.PorterAcceptJob(porter)
	assert.Equal(0, len(pool.Porters), "pool should be empty")
	assert.Equal(domain.PorterStatusWorking, porter.Status, "porter status should be working")
}

func TestPorterUnavailable(t *testing.T) {
	assert := assert.New(t)

	pool := domain.CreateNewPool()
	porter, err := domain.CreatedNewPorter("John Smith", "porter-001")
	assert.Nil(err, "error should be nil")
	pool.PorterRegister(porter, "token-001")
	pool.PorterUnavailable(porter)
	assert.Equal(0, len(pool.Porters), "pool should be empty")
	assert.Equal(domain.PorterStatusUnavailable, porter.Status, "porter status should be unavailable")
}
