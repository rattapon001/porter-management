package main

import (
	"github.com/gin-gonic/gin"
	job_router "github.com/rattapon001/porter-management/api/v1/routers/job"
	porter_router "github.com/rattapon001/porter-management/api/v1/routers/porter"
)

func main() {
	port := ":8080"
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	job_router.InitJobRouter(router)
	porter_router.InitPorterRouter(router)
	router.Run(port)
}
