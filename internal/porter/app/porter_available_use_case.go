package app

import (
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

func (s *PorterUseCaseImpl) PorterAvailable(code domain.PorterCode) (*domain.Porter, error) {
	porter, err := s.Repo.FindByCode(code)
	if err != nil {
		return nil, err
	}
	porter.Available()
	if err := s.Repo.Save(porter); err != nil {
		return nil, err
	}
	return porter, nil
}
