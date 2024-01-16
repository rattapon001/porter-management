package job_router

import (
	"github.com/gin-gonic/gin"
	job_handler "github.com/rattapon001/porter-management/api/v1/handlers/job"
	"github.com/rattapon001/porter-management/internal/job/app"
)

func InitJobRouter(router *gin.Engine, JobService app.JobService) {
	JobHandler := job_handler.NewJobHandler(JobService)

	jobRouter := router.Group("/jobs")
	{
		jobRouter.POST("/", JobHandler.CreatedNewJob)
		jobRouter.POST("/:id/accepted", JobHandler.AcceptedJob)
		jobRouter.POST("/:id/started", JobHandler.StartedJob)
		jobRouter.POST("/:id/completed", JobHandler.CompletedJob)
		jobRouter.GET("/:id", JobHandler.FindJobById)
	}
}
