package main

import (
	"github.com/gin-gonic/gin"
	job_router "github.com/rattapon001/porter-management/api/v1/routers/job"
	porter_router "github.com/rattapon001/porter-management/api/v1/routers/porter"
	postgresorm "github.com/rattapon001/porter-management/internal/infra/postgres_orm"
	job_app "github.com/rattapon001/porter-management/internal/job/app"
	job_domain "github.com/rattapon001/porter-management/internal/job/domain"
	job_memory "github.com/rattapon001/porter-management/internal/job/infra/memory"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	porter_app "github.com/rattapon001/porter-management/internal/porter/app"
	porter_domain "github.com/rattapon001/porter-management/internal/porter/domain"
	porter_memory "github.com/rattapon001/porter-management/internal/porter/infra/memory"
	porter_postgres "github.com/rattapon001/porter-management/internal/porter/infra/postgres_orm"
)

func main() {

	dbConfig := postgresorm.GetPostgresOrmConfig()
	db, err := postgresorm.NewPostgresOrmDbConnection(dbConfig)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&job_domain.Job{}, &porter_domain.Porter{})

	port := ":8080"
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Init Job Router
	jobRepository := job_postgres.NewPostgresOrmRepository(db)
	publisher := job_memory.NewMockImplimentEventHandler()
	JobUseCase := job_app.NewJobUseCase(jobRepository, publisher)
	job_router.InitJobRouter(router, JobUseCase)

	// Init Porter Router
	porterRepository := porter_postgres.NewPostgresOrmRepository(db)
	PorterPublisher := porter_memory.NewMockImplimentEventHandler()
	PorterUseCase := porter_app.NewPorterUseCase(porterRepository, PorterPublisher)
	porter_router.InitPorterRouter(router, PorterUseCase)
	router.Run(port)
}
