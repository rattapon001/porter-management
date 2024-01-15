package memory

import (
	"fmt"

	"github.com/rattapon001/porter-management/pkg"
)

type MockImplimentEventHandler struct {
}

func NewMockImplimentEventHandler() *MockImplimentEventHandler {
	return &MockImplimentEventHandler{}
}

func (h *MockImplimentEventHandler) Publish(event []pkg.Event) error {
	fmt.Printf("Publish event: %v\n", event[len(event)-1])
	return nil
}
