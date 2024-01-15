package job_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *JobHandler) CreatedNewJob(c *gin.Context) {
	var job domain.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	location := domain.Location{
		From: job.Location.From,
		To:   job.Location.To,
	}
	patient := domain.Patient{
		Name: job.Patient.Name,
		HN:   job.Patient.HN,
	}

	newJob, err := h.JobService.CreatedNewJob(location, patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newJob)
}

func (h *JobHandler) AcceptedJob(c *gin.Context) {
	var porter domain.Porter
	if err := c.ShouldBindJSON(&porter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	jobId := domain.JobId(id)

	job, err := h.JobService.AcceptedJob(jobId, porter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) FindJobById(c *gin.Context) {
	id := c.Param("id")
	jobId := domain.JobId(id)

	job, err := h.JobService.FindJobById(jobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}
