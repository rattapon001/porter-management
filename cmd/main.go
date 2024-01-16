package main

import (
	"github.com/gin-gonic/gin"
	job_router "github.com/rattapon001/porter-management/api/v1/routers/job"
	porter_router "github.com/rattapon001/porter-management/api/v1/routers/porter"
	job_app "github.com/rattapon001/porter-management/internal/job/app"
	job_memory "github.com/rattapon001/porter-management/internal/job/infra/memory"
	porter_app "github.com/rattapon001/porter-management/internal/porter/app"
	porter_memory "github.com/rattapon001/porter-management/internal/porter/infra/memory"
)

func main() {
	port := ":8080"
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Init Job Router
	jobRepository := job_memory.NewJobMemoryRepository()
	publisher := job_memory.NewMockImplimentEventHandler()
	JobUseCase := job_app.NewJobUseCase(jobRepository, publisher)
	job_router.InitJobRouter(router, JobUseCase)

	// Init Porter Router
	porterRepository := porter_memory.NewPorterMemoryRepository()
	PorterPublisher := porter_memory.NewMockImplimentEventHandler()
	PorterUseCase := porter_app.NewPorterUseCase(porterRepository, PorterPublisher)
	porter_router.InitPorterRouter(router, PorterUseCase)
	router.Run(port)
}
