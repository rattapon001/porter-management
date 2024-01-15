package job_handler

import (
	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type JobHandler struct {
	JobService app.JobService
}

func NewJobHandler(jobService app.JobService) *JobHandler {
	return &JobHandler{
		JobService: jobService,
	}
}

func (h *JobHandler) CreatedNewJob(location domain.Location, patient domain.Patient) error {
	_, err := h.JobService.CreatedNewJob(location, patient)
	if err != nil {
		return err
	}
	return nil
}

func (h *JobHandler) AcceptedJob(id string, porter domain.Porter) error {
	_, err := h.JobService.AcceptedJob(id, porter)
	if err != nil {
		return err
	}
	return nil
}
