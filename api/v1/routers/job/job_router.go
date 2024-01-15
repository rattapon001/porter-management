package job_router

import (
	"github.com/gin-gonic/gin"
	job_handler "github.com/rattapon001/porter-management/api/v1/handlers/job"
	"github.com/rattapon001/porter-management/internal/job/app"
	"github.com/rattapon001/porter-management/internal/job/infra/memory"
)

func InitJobRouter(router *gin.Engine) {
	jobRepository := memory.NewJobMemoryRepository()
	publisher := memory.NewMockImplimentEventHandler()
	JobService := app.NewJobService(jobRepository, publisher)
	JobHandler := job_handler.NewJobHandler(JobService)

	jobRouter := router.Group("/jobs")
	{
		jobRouter.POST("/", JobHandler.CreatedNewJob)
		jobRouter.POST("/:id/accepted", JobHandler.AcceptedJob)
		jobRouter.GET("/:id", JobHandler.FindJobById)
	}
}
