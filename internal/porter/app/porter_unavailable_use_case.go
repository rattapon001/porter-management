package app

import (
	"github.com/rattapon001/porter-management/internal/porter/domain"
	error_handler "github.com/rattapon001/porter-management/internal/porter/domain/errors"
)

func (s *PorterServiceImpl) PorterUnavailable(ID domain.PorterId) (*domain.Porter, error) {
	porter := s.Repo.FindByID(ID)
	if porter == nil {
		return nil, error_handler.ErrPorterNotFound
	}
	porter.Unavailable()
	if err := s.Repo.Update(porter); err != nil {
		return nil, err
	}
	return porter, nil
}
