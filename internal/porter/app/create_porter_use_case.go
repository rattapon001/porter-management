package app

import (
	"github.com/google/uuid"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

func (s *PorterUseCaseImpl) CreateNewPorter(name string, token string) (*domain.Porter, error) {
	portCode, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	porter, err := domain.NewPorter(name, portCode.String(), token)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Save(porter)
	if err != nil {
		return nil, err
	}
	return porter, nil
}
