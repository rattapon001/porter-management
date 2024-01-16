package porter_router

import (
	"github.com/gin-gonic/gin"
	porter_handler "github.com/rattapon001/porter-management/api/v1/handlers/porter"
	"github.com/rattapon001/porter-management/internal/porter/app"
)

func InitPorterRouter(router *gin.Engine, porterService app.PorterService) {
	porterHandler := porter_handler.NewPorterHandler(porterService)
	porterRouter := router.Group("/porters")
	{
		porterRouter.POST("/", porterHandler.CreatedNewPorter)
		porterRouter.PUT("/:code/available", porterHandler.PorterAvailable)
	}
}
