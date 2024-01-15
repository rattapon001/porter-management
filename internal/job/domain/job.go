package domain

import (
	"time"

	"github.com/google/uuid"
)

type JobId string
type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusAccepted  JobStatus = "accepted"
	JobStatusWorking   JobStatus = "working"
	JobStatusCompleted JobStatus = "completed"
)

type Location struct {
	From string `bson:"from"`
	To   string `bson:"to"`
}

type Patient struct {
	Name string `bson:"name"`
	HN   string `bson:"hn"`
}

type Porter struct {
	Code string `bson:"code"`
	Name string `bson:"name"`
}

type Job struct {
	ID        JobId     `bson:"_id,omitempty"`
	Version   int       `bson:"version"`
	Status    JobStatus `bson:"status"`
	Accepted  bool      `bson:"accepted"`
	Location  Location  `bson:"location"`
	Patient   Patient   `bson:"patient"`
	Porter    Porter    `bson:"porter"`
	CheckIn   time.Time `bson:"check_in"`
	CheckOut  time.Time `bson:"check_out"`
	Aggregate Aggregate `bson:"aggregate"`
}

func CreatedNewJob(location Location, patient Patient) (*Job, error) {

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	job := &Job{
		ID:       JobId(ID.String()),
		Version:  1,
		Status:   JobStatusPending,
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
	j.Aggregate.AppendEvent(JobEventCreated, payload)
}

func (j *Job) AcceptedJob(porter Porter) {
	if j.Status != JobStatusPending {
		return
	}
	j.Status = JobStatusAccepted
	j.Accepted = true
	j.Porter = porter
	j.JobAcceptedEvent()
}

func (j *Job) StartedJob() {
	if j.Status != JobStatusAccepted {
		return
	}
	j.CheckInJob()
	j.Status = JobStatusWorking
	j.JobStartedEvent()
}

func (j *Job) CompletedJob() {
	if j.Status != JobStatusWorking {
		return
	}
	j.CheckOutJob()
	j.Status = JobStatusCompleted
	j.JobCompletedEvent()
}

func (j *Job) CheckInJob() {
	j.CheckIn = time.Now()
}

func (j *Job) CheckOutJob() {
	j.CheckOut = time.Now()
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
	j.Aggregate.AppendEvent(JobEventAccepted, payload)
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
	j.Aggregate.AppendEvent(JobEventWorking, payload)
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
	j.Aggregate.AppendEvent(JobEventCompleted, payload)
}
