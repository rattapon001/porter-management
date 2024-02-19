package stock_router

import (
	"github.com/gin-gonic/gin"
	stock_handler "github.com/rattapon001/porter-management/api/v1/handlers/stock"
	"github.com/rattapon001/porter-management/internal/stock/app"
)

func InitStockRouter(router *gin.Engine, StockUseCase app.StockUseCase) {
	stockHandler := stock_handler.NewStockHandler(StockUseCase)
	stockRouter := router.Group("/stock")
	{
		stockRouter.POST("/", stockHandler.CreateItem)
	}
}
