package stock_handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/porter-management/internal/stock/app"
	"github.com/rattapon001/porter-management/internal/stock/domain"
)

type StockHandler struct {
	StockUseCase app.StockUseCase
}

func NewStockHandler(StockUseCase app.StockUseCase) *StockHandler {
	return &StockHandler{
		StockUseCase: StockUseCase,
	}
}

func (h *StockHandler) CreateItem(c *gin.Context) {
	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem, err := h.StockUseCase.CreateItem(context.Background(), &item) // Pass a pointer to the item variable
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newItem)
}
