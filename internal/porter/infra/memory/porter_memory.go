package memory

import "github.com/rattapon001/porter-management/internal/porter/domain"

type PorterMemoryRepository struct {
	porters []*domain.Porter
}

func NewPorterMemoryRepository() *PorterMemoryRepository {
	return &PorterMemoryRepository{
		porters: []*domain.Porter{},
	}
}

func (r *PorterMemoryRepository) Save(porter *domain.Porter) error {
	r.porters = append(r.porters, porter)
	return nil
}
