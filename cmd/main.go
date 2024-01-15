package main

import (
	"github.com/gin-gonic/gin"
	job_router "github.com/rattapon001/porter-management/api/v1/routers/job"
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
	router.Run(port)
}
