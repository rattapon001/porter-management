package porter_router

import (
	"github.com/gin-gonic/gin"
	porter_handler "github.com/rattapon001/porter-management/api/v1/handlers/porter"
	"github.com/rattapon001/porter-management/internal/porter/app"
)

func InitPorterRouter(router *gin.Engine, PorterUseCase app.PorterUseCase) {
	porterHandler := porter_handler.NewPorterHandler(PorterUseCase)
	porterRouter := router.Group("/porters")
	{
		porterRouter.POST("/", porterHandler.NewPorter)
		porterRouter.PUT("/:code/available", porterHandler.PorterAvailable)
	}
}
