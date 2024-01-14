package app

import (
	"github.com/rattapon001/porter-management/internal/porter/domain"
	error_handler "github.com/rattapon001/porter-management/internal/porter/domain/errors"
)

func (s *PorterServiceImpl) ReadyForJob(ID domain.PorterId) error {
	porter := s.Repo.FindByID(ID)
	if porter == nil {
		return error_handler.ErrPorterNotFound
	}
	porter.Available()
	err := s.Repo.Update(porter)
	if err != nil {
		return err
	}
	return nil
}
