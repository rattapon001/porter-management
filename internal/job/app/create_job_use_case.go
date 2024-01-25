package app

import (
	"log"

	"github.com/google/uuid"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

func (s *JobUseCaseImpl) CreateNewJob(location domain.Location, patient domain.Patient) (*domain.Job, error) {

	eId := []domain.EquipmentId{domain.EquipmentId(uuid.New().String()), domain.EquipmentId(uuid.New().String())}
	job, err := domain.NewJob(location, patient, eId)
	log.Printf("job: %+v", job)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Save(job)
	if err != nil {
		return nil, err
	}
	if err := s.Publisher.Publish(job.Aggregate.Events); err != nil {
		return nil, err
	}
	return job, nil
}
