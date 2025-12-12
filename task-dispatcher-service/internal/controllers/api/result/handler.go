package result

import (
	"net/http"
	"task-dispatcher-service/internal/domain/dto"
	"task-dispatcher-service/internal/domain/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResultHandler struct {
	logger  *zap.Logger
	service *services.TaskService
}

func NewResultHandler(logger *zap.Logger, rg *gin.RouterGroup, service *services.TaskService) *ResultHandler {
	handler := &ResultHandler{
		logger:  logger,
		service: service,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (h *ResultHandler) ProcessTask(c *gin.Context) {
	var task dto.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go h.service.HandleResult(task)

	c.JSON(http.StatusOK, "ok")
}

func (h *ResultHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("", h.ProcessTask)
}
