package app

import (
	"github.com/rattapon001/porter-management/internal/porter/domain"
	error_handler "github.com/rattapon001/porter-management/internal/porter/domain/errors"
)

func (s *PorterUseCaseImpl) PorterWorking(code domain.PorterCode) (*domain.Porter, error) {
	porter := s.Repo.FindByCode(code)
	if porter == nil {
		return nil, error_handler.ErrPorterNotFound
	}
	porter.Working()
	if err := s.Repo.Update(porter); err != nil {
		return nil, err
	}
	return porter, nil
}
