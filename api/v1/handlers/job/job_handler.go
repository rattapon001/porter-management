package job_handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/domain"
)

type JobHandler struct {
	JobUseCase app.JobUseCase
}

func NewJobHandler(JobUseCase app.JobUseCase) *JobHandler {
	return &JobHandler{
		JobUseCase: JobUseCase,
	}
}

func (h *JobHandler) CreateNewJob(c *gin.Context) {
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
	equipments := []domain.Equipment{}
	for _, e := range job.Equipments {
		equipment := domain.Equipment{
			EquipmentId: e.EquipmentId,
			Name:        e.Name,
			Amount:      e.Amount,
		}
		equipments = append(equipments, equipment)
	}

	newJob, err := h.JobUseCase.CreateNewJob(context.Background(), location, patient, equipments)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newJob)
}

func (h *JobHandler) AcceptJob(c *gin.Context) {
	var porter domain.Porter
	if err := c.ShouldBindJSON(&porter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	jobId := domain.JobId(id)

	job, err := h.JobUseCase.AcceptJob(jobId, porter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) StartJob(c *gin.Context) {
	id := c.Param("id")
	jobId := domain.JobId(id)

	job, err := h.JobUseCase.StartJob(jobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) CompleteJob(c *gin.Context) {
	id := c.Param("id")
	jobId := domain.JobId(id)
	job, err := h.JobUseCase.CompleteJob(jobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}

func (h *JobHandler) FindJobById(c *gin.Context) {
	id := c.Param("id")
	jobId := domain.JobId(id)

	job, err := h.JobUseCase.FindJobById(jobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}
