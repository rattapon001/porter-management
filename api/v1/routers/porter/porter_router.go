package porter_router

import (
	"github.com/gin-gonic/gin"
	porter_handler "github.com/rattapon001/porter-management/api/v1/handlers/porter"
	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/infra/memory"
)

func InitPorterRouter(router *gin.Engine) {

	porterRepository := memory.NewPorterMemoryRepository()
	publisher := memory.NewMockImplimentEventHandler()
	porterService := app.NewPorterService(porterRepository, publisher)
	porterHandler := porter_handler.NewPorterHandler(porterService)

	porterRouter := router.Group("/porters")
	{
		porterRouter.POST("/", porterHandler.CreatedNewPorter)
		porterRouter.PUT("/:code/available", porterHandler.PorterAvailable)
	}
}
