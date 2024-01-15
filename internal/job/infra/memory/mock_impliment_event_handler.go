package memory

import "github.com/rattapon001/porter-management/pkg"

type MockImplimentEventHandler struct {
}

func NewMockImplimentEventHandler() *MockImplimentEventHandler {
	return &MockImplimentEventHandler{}
}

func (h *MockImplimentEventHandler) Publish(event []pkg.Event) error {
	return nil
}
