package process

import (
	"net/http"
	"stub-file-handler/internal/domain/dto"
	"stub-file-handler/internal/domain/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProcessTaskHandler struct {
	logger  *zap.Logger
	service *services.StubService
}

func NewProcessTaskHandler(logger *zap.Logger, rg *gin.RouterGroup, service *services.StubService) *ProcessTaskHandler {
	handler := &ProcessTaskHandler{
		logger:  logger,
		service: service,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (h *ProcessTaskHandler) ProcessTask(c *gin.Context) {
	var task dto.StubFileTask

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.logger.Sugar().Info(task)
	go h.service.ProcessTask(task)

	c.JSON(http.StatusOK, "ok")
}

func (h *ProcessTaskHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("", h.ProcessTask)
}
