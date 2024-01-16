package porter_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rattapon001/porter-management/internal/porter/app"
	"github.com/rattapon001/porter-management/internal/porter/domain"
)

type PorterHandler struct {
	PorterService app.PorterService
}

func NewPorterHandler(porterService app.PorterService) *PorterHandler {
	return &PorterHandler{
		PorterService: porterService,
	}
}

func (h *PorterHandler) NewPorter(c *gin.Context) {
	var porter domain.Porter
	if err := c.ShouldBindJSON(&porter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mockNotiToken := uuid.New().String()
	newPorter, err := h.PorterService.CreateNewPorter(porter.Name, mockNotiToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newPorter)
}

func (h *PorterHandler) PorterAvailable(c *gin.Context) {
	code := domain.PorterCode(c.Param("code"))
	porter, err := h.PorterService.PorterAvailable(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, porter)
}

func (h *PorterHandler) PorterUnavailable(c *gin.Context) {
	code := domain.PorterCode(c.Param("code"))
	porter, err := h.PorterService.PorterUnavailable(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, porter)
}
