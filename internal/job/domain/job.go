package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

type Location struct {
	From string
	To   string
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *Location) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &l)

		return err
	}
	return fmt.Errorf("failed to unmarshal Location value: %v", value)
}

type Patient struct {
	Name string
	HN   string
}

func (p Patient) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Patient) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &p)
		return err
	}
	return fmt.Errorf("failed to unmarshal PatientDB value: %v", value)
}

type Porter struct {
	Code string
	Name string
}

func (p Porter) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Porter) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &p)
		return err
	}
	return fmt.Errorf("failed to unmarshal PorterDB value: %v", value)
}

type Job struct {
	ID        JobId `gorm:"primaryKey"`
	Version   int
	Status    JobStatus
	Accepted  bool
	Location  Location `gorm:"type:jsonb"`
	Patient   Patient  `gorm:"type:jsonb"`
	Porter    Porter   `gorm:"type:jsonb"`
	CheckIn   time.Time
	CheckOut  time.Time
	Aggregate Aggregate `gorm:"type:jsonb"`
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
