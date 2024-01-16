package job_router

import (
	"github.com/gin-gonic/gin"
	job_handler "github.com/rattapon001/porter-management/api/v1/handlers/job"
	"github.com/rattapon001/porter-management/internal/job/app"
)

func InitJobRouter(router *gin.Engine, JobUseCase app.JobUseCase) {
	JobHandler := job_handler.NewJobHandler(JobUseCase)

	jobRouter := router.Group("/jobs")
	{
		jobRouter.POST("/", JobHandler.CreateNewJob)
		jobRouter.POST("/:id/accept", JobHandler.AcceptJob)
		jobRouter.POST("/:id/start", JobHandler.StartJob)
		jobRouter.POST("/:id/complete", JobHandler.CompleteJob)
		jobRouter.GET("/:id", JobHandler.FindJobById)
	}
}
