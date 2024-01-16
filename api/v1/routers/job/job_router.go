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
		jobRouter.POST("/", JobHandler.CreateNewJob)
		jobRouter.POST("/:id/accept", JobHandler.AcceptJob)
		jobRouter.POST("/:id/start", JobHandler.StartJob)
		jobRouter.POST("/:id/complete", JobHandler.CompleteJob)
		jobRouter.GET("/:id", JobHandler.FindJobById)
	}
}
