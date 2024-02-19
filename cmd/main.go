package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	job_router "github.com/rattapon001/porter-management/api/v1/routers/job"
	porter_router "github.com/rattapon001/porter-management/api/v1/routers/porter"
	stock_router "github.com/rattapon001/porter-management/api/v1/routers/stock"
	"github.com/rattapon001/porter-management/internal/infra/kafka"
	postgresorm "github.com/rattapon001/porter-management/internal/infra/postgres_orm"
	"github.com/rattapon001/porter-management/internal/infra/uow"
	job_app "github.com/rattapon001/porter-management/internal/job/app"
	job_postgres "github.com/rattapon001/porter-management/internal/job/infra/postgres_orm"
	porter_app "github.com/rattapon001/porter-management/internal/porter/app"
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

	kafkaConsumer, err := kafka.InitKafkaConsumer()
	if err != nil {
		panic(err)
	}

	uow := uow.NewUnitOfWork(db)

	// Init Job Router
	jobRepository := job_postgres.NewPostgresOrmRepository(db)
	publisher, err := kafka.NewProducer()
	if err != nil {
		panic(err)
	}
	JobUseCase := job_app.NewJobUseCase(jobRepository, publisher, uow)
	job_router.InitJobRouter(router, JobUseCase)

	// Init Porter Router
	porterRepository := porter_postgres.NewPostgresOrmRepository(db)
	PorterPublisher := porter_memory.NewMockImplimentEventHandler()
	PorterUseCase := porter_app.NewPorterUseCase(porterRepository, PorterPublisher)
	porter_router.InitPorterRouter(router, PorterUseCase)

	stockRepository := stock_postgres.NewPostgresOrmRepository(db)
	stockPublisher, err := kafka.NewProducer()
	if err != nil {
		panic(err)
	}
	stockUseCase := stock_app.NewStockUseCase(stockRepository, stockPublisher, uow)

	stockConsumer := stock_kafka.NewKafkaConsumer(kafkaConsumer, stockUseCase)
	stockConsumer.Subscribe([]string{"job_created"})
	stock_router.InitStockRouter(router, stockUseCase)

	// Start the server
	router.Run(port)
}
