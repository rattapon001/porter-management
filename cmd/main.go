package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	job_router "github.com/rattapon001/porter-management/api/v1/routers/job"
	porter_router "github.com/rattapon001/porter-management/api/v1/routers/porter"
	stock_router "github.com/rattapon001/porter-management/api/v1/routers/stock"
	"github.com/rattapon001/porter-management/internal/infra/kafka"
	"github.com/rattapon001/porter-management/internal/infra/outbox/publisher"
	postgresorm "github.com/rattapon001/porter-management/internal/infra/postgres_orm"
	"github.com/rattapon001/porter-management/internal/infra/uow"
	job_app "github.com/rattapon001/porter-management/internal/job/app"
	job_domain "github.com/rattapon001/porter-management/internal/job/domain"
	job_kafka "github.com/rattapon001/porter-management/internal/job/infra/kafka"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	porter_app "github.com/rattapon001/porter-management/internal/porter/app"
	porter_kafka "github.com/rattapon001/porter-management/internal/porter/infra/kafka"
	porter_memory "github.com/rattapon001/porter-management/internal/porter/infra/memory"
	porter_postgres "github.com/rattapon001/porter-management/internal/porter/infra/postgres_orm"
	stock_app "github.com/rattapon001/porter-management/internal/stock/app"
	stock_kafka "github.com/rattapon001/porter-management/internal/stock/infra/kafka"
	stock_postgres "github.com/rattapon001/porter-management/internal/stock/infra/postgres_orm"
)

func main() {

	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbConfig := postgresorm.GetPostgresOrmConfig()
	db, err := postgresorm.NewPostgresOrmDbConnection(dbConfig)

	if err != nil {
		panic(err)
	}

	port := ":8080"
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	uow := uow.NewUnitOfWork(db)

	// Init Job Router
	jobRepository := job_postgres.NewPostgresOrmRepository(db)
	publisher := publisher.NewPublisher()

	initJobConsumer, err := kafka.InitKafkaConsumer()
	if err != nil {
		panic(err)
	}

	jobDql, err := kafka.NewKafkaDLQs()
	if err != nil {
		panic(err)
	}

	JobUseCase := job_app.NewJobUseCase(jobRepository, publisher, uow)
	jobComsumer := job_kafka.NewKafkaConsumer(initJobConsumer, JobUseCase, jobDql)
	jobComsumer.Subscribe([]string{string("job.events." + job_domain.ItemAllocatedEvent)})
	job_router.InitJobRouter(router, JobUseCase)

	// Init Porter Router

	initPorterConsumer, err := kafka.InitKafkaConsumer()
	if err != nil {
		panic(err)
	}

	porterDql, err := kafka.NewKafkaDLQs()
	if err != nil {
		panic(err)
	}

	porterRepository := porter_postgres.NewPostgresOrmRepository(db)
	PorterPublisher := porter_memory.NewMockImplimentEventHandler()
	PorterUseCase := porter_app.NewPorterUseCase(porterRepository, PorterPublisher)
	porterConsumer := porter_kafka.NewKafkaConsumer(initPorterConsumer, PorterUseCase, porterDql)
	porterConsumer.Subscribe([]string{"job.events." + string(job_domain.JobAllocatedEvent)})
	porter_router.InitPorterRouter(router, PorterUseCase)

	stockRepository := stock_postgres.NewPostgresOrmRepository(db)

	initStockConsumer, err := kafka.InitKafkaConsumer()
	if err != nil {
		panic(err)
	}

	stockDql, err := kafka.NewKafkaDLQs()
	if err != nil {
		panic(err)
	}

	stockUseCase := stock_app.NewStockUseCase(stockRepository, publisher, uow)
	stockConsumer := stock_kafka.NewKafkaConsumer(initStockConsumer, stockUseCase, stockDql)
	stockConsumer.Subscribe([]string{"job.events." + string(job_domain.JobCreatedEvent)})
	stock_router.InitStockRouter(router, stockUseCase)

	// Start the server
	router.Run(port)
}
