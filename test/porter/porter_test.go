package porter_test

import (
	"testing"

	"github.com/rattapon001/porter-management/internal/porter/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewPorter(t *testing.T) {
	assert := assert.New(t)

	createdPorter, _ := domain.CreateNewPorter("John Smith", "porter-001")
	assert.Equal(domain.PorterStatusUnavailable, createdPorter.Status, "created porter status should be unavailable")
}
