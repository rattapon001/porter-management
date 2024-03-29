package domain

type JobId string

type Location struct {
	From string `bson:"from"`
	To   string `bson:"to"`
}

type Patient struct {
	Name string `bson:"name"`
	HN   string `bson:"hn"`
}

type Job struct {
	ID       JobId
	Version  int
	Patient  Patient
	Location Location
	Porter   Porter
}

func NewJob(id JobId, patient Patient, location Location) *Job {
	return &Job{
		ID:       id,
		Patient:  patient,
		Location: location,
	}
}
