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

func (r *PorterMemoryRepository) Update(porter *domain.Porter) error {
	for i, p := range r.porters {
		if p.ID == porter.ID {
			r.porters[i] = porter
			return nil
		}
	}
	return nil
}

func (r *PorterMemoryRepository) FindAvailablePorter() *domain.Porter {
	for _, p := range r.porters {
		if p.Status == domain.PorterStatusAvailable {
			return p
		}
	}
	return nil
}

func (r *PorterMemoryRepository) FindByID(id domain.PorterId) *domain.Porter {
	for _, p := range r.porters {
		if p.ID == id {
			return p
		}
	}
	return nil
}

func (r *PorterMemoryRepository) FindByCode(code string) *domain.Porter {
	for _, p := range r.porters {
		if p.PorterCode == code {
			return p
		}
	}
	return nil
}
