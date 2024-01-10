package domain

type JobRepository interface {
	Save(job *Job) error
	Update(job *Job) error
	FindById(id JobId) (*Job, error)
	FindAll() ([]*Job, error)
	Delete(job *Job) error
}
