package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewPorter(t *testing.T) {
	assert := assert.New(t)

	createdPorter, err := domain.NewPorter("John Smith", "porter-001", "token")
	assert.Nil(err, "error should be nil")
	assert.Equal(domain.PorterStatusUnavailable, createdPorter.Status, "created porter status should be unavailable")
}
