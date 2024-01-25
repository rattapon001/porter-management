package domain

import (
	"time"

	"github.com/google/uuid"
	domain_errors "github.com/rattapon001/porter-management/internal/job/domain/errors"
)

type JobId string     // JobId is a unique identifier for a job
type JobStatus string // JobStatus is a status of a job

const (
	JobPendingStatus   JobStatus = "pending"   // JobPendingStatus is a status of a job when it is created
	JobAcceptedStatus  JobStatus = "accepted"  // JobAcceptedStatus is a status of a job when it is accepted by a porter
	JobWorkingStatus   JobStatus = "working"   // JobWorkingStatus is a status of a job when it is started
	JobCompletedStatus JobStatus = "completed" // JobCompletedStatus is a status of a job when it is completed
)

type Job struct {
	ID         JobId       `bson:"_id" gorm:"primaryKey;type:uuid"`
	Version    int         `bson:"version"`
	Status     JobStatus   `bson:"status"`
	Accepted   bool        `bson:"accepted"`
	Location   Location    `bson:"location" gorm:"type:jsonb"`
	Patient    Patient     `bson:"patient" gorm:"type:jsonb"`
	Porter     Porter      `bson:"porter" gorm:"type:jsonb"`
	CheckIn    time.Time   `bson:"checkIn"`
	CheckOut   time.Time   `bson:"checkOut"`
	Equipments []Equipment `bson:"equipments" gorm:"foreignKey:JobId"`
	Aggregate  Aggregate   `bson:"aggregate" gorm:"type:jsonb"`
}

func NewJob(location Location, patient Patient, equipmentIds []EquipmentId) (*Job, error) {
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
	for _, eId := range equipmentIds {
		equipment, _ := NewEquipment(eId, job.ID, "test", 1)
		job.AddEquipment(*equipment)
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
		return domain_errors.ErrCannotAcceptJob
	}
	j.Status = JobAcceptedStatus
	j.Accepted = true
	j.Porter = porter
	j.JobAcceptedEvent()
	return nil
}

func (j *Job) Start() error {
	if j.Status != JobAcceptedStatus {
		return domain_errors.ErrCannotStartJob
	}
	j.CheckIn = time.Now()
	j.Status = JobWorkingStatus
	j.JobStartedEvent()
	return nil
}

func (j *Job) Complete() error {
	if j.Status != JobWorkingStatus {
		return domain_errors.ErrCannotCompleteJob
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

func (j *Job) AddEquipment(equipment Equipment) error {
	if j.Status != JobPendingStatus {
		return domain_errors.ErrCannotAddEquipment
	}
	j.Equipments = append(j.Equipments, equipment)
	return nil
}
