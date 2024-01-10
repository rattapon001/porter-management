package domain

type JobRepository interface {
	Save(job *Job) error
}
