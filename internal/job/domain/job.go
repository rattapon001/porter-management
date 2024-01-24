package domain

import (
	"time"

	"github.com/google/uuid"
	domain_errors "github.com/rattapon001/porter-management/internal/job/domain/errors"
)

type JobId string
type JobStatus string

const (
	JobPendingStatus   JobStatus = "pending"
	JobAcceptedStatus  JobStatus = "accepted"
	JobWorkingStatus   JobStatus = "working"
	JobCompletedStatus JobStatus = "completed"
)

type Job struct {
	ID        JobId     `bson:"_id" gorm:"primaryKey"`
	Version   int       `bson:"version"`
	Status    JobStatus `bson:"status"`
	Accepted  bool      `bson:"accepted"`
	Location  Location  `bson:"location" gorm:"type:jsonb"`
	Patient   Patient   `bson:"patient" gorm:"type:jsonb"`
	Porter    Porter    `bson:"porter" gorm:"type:jsonb"`
	CheckIn   time.Time `bson:"checkIn"`
	CheckOut  time.Time `bson:"checkOut"`
	Aggregate Aggregate `bson:"aggregate" gorm:"type:jsonb"`
}

func NewJob(location Location, patient Patient) (*Job, error) {

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	job := &Job{
		ID:       JobId(ID.String()),
		Version:  1,
		Status:   JobPendingStatus,
		Location: location,
		Patient:  patient,
	}
	job.JobCreatedEvent()
	return job, nil
}

func (j *Job) JobCreatedEvent() {

	payload := map[string]interface{}{
		"job_id":   j.ID,
		"version":  j.Version,
		"status":   j.Status,
		"location": j.Location,
		"patient":  j.Patient,
	}
	j.Aggregate.AppendEvent(JobCreatedEvent, payload)
}

func (j *Job) Accept(porter Porter) error {
	if j.Status != JobPendingStatus {
		return domain_errors.CannotAcceptJob
	}
	j.Status = JobAcceptedStatus
	j.Accepted = true
	j.Porter = porter
	j.JobAcceptedEvent()
	return nil
}

func (j *Job) Start() error {
	if j.Status != JobAcceptedStatus {
		return domain_errors.CannotStartJob
	}
	j.CheckIn = time.Now()
	j.Status = JobWorkingStatus
	j.JobStartedEvent()
	return nil
}

func (j *Job) Complete() error {
	if j.Status != JobWorkingStatus {
		return domain_errors.CannotCompleteJob
	}
	j.CheckOut = time.Now()
	j.Status = JobCompletedStatus
	j.JobCompletedEvent()
	return nil
}

func (j *Job) JobAcceptedEvent() {

	payload := map[string]interface{}{
		"job_id":   j.ID,
		"version":  j.Version + 1,
		"status":   j.Status,
		"location": j.Location,
		"patient":  j.Patient,
		"porter":   j.Porter,
	}
	j.Aggregate.AppendEvent(JobAcceptedEvent, payload)
}

func (j *Job) JobStartedEvent() {

	payload := map[string]interface{}{
		"job_id":   j.ID,
		"version":  j.Version + 1,
		"status":   j.Status,
		"location": j.Location,
		"patient":  j.Patient,
		"porter":   j.Porter,
		"check_in": j.CheckIn,
	}
	j.Aggregate.AppendEvent(JobWorkingEvent, payload)
}

func (j *Job) JobCompletedEvent() {
	payload := map[string]interface{}{
		"job_id":    j.ID,
		"version":   j.Version + 1,
		"status":    j.Status,
		"location":  j.Location,
		"patient":   j.Patient,
		"porter":    j.Porter,
		"check_in":  j.CheckIn,
		"check_out": j.CheckOut,
	}
	j.Aggregate.AppendEvent(JobCompletedEvent, payload)
}
