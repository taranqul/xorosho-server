package worker

import (
	"net/http"
	"worker-manager-service/internal/domain/dto"
	"worker-manager-service/internal/domain/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WorkerHandler struct {
	logger         *zap.Logger
	worker_service *service.WorkerService
}

func NewWorkerHandler(service *service.WorkerService, logger *zap.Logger, rg *gin.RouterGroup) *WorkerHandler {
	handler := &WorkerHandler{
		logger:         logger,
		worker_service: service,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (w *WorkerHandler) RegisterWorker(c *gin.Context) {
	var worker dto.WorkerRegister

	if err := c.BindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := w.worker_service.RegisterService(worker)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (*WorkerHandler) ValidateScheme(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func (h *WorkerHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("", h.RegisterWorker)
	rg.POST("/validate", h.ValidateScheme)
}
