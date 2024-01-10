package domain

import (
	"time"

	"github.com/google/uuid"
)

type JobId string
type JobStatus string

const (
	JobStatusPending JobStatus = "pending"
	Accepted         JobStatus = "accepted"
	Working          JobStatus = "working"
	Completed        JobStatus = "completed"
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

func CreateNewJob(location Location, patient Patient) (*Job, error) {

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
	event := Event{
		EventName: JobCreated,
		Payload: map[string]interface{}{
			"job_id":   j.ID,
			"version":  j.Version,
			"status":   j.Status,
			"location": j.Location,
			"patient":  j.Patient,
		},
	}
	j.Aggregate.AppendEvent(JobCreated, event)
}
