package app

import "github.com/rattapon001/porter-management/internal/porter/domain"

func (s *PorterServiceImpl) ReadyForJob(porter *domain.Porter, token string) error {
	s.Pool.PorterRegister(porter, token)
	s.Repo.Save(porter)
	return nil
}
