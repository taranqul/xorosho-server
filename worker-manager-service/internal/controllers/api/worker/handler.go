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
	w.logger.Sugar().Debug("registration triggered")
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

func (w *WorkerHandler) ValidateScheme(c *gin.Context) {
	var worker_scheme dto.TaskRequest

	if err := c.BindJSON(&worker_scheme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := w.worker_service.ValidateScheme(worker_scheme)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !res {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

func (w *WorkerHandler) GetScheme(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter 'name' is required",
		})
		return
	}
	scheme, err := w.worker_service.GetScheme(name)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, scheme)
}

func (w *WorkerHandler) GetTasks(c *gin.Context) {
	list, err := w.worker_service.GetTasks()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *WorkerHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("", h.RegisterWorker)
	rg.POST("/validate", h.ValidateScheme)
	rg.GET("", h.GetTasks)
	rg.GET("/scheme", h.GetScheme)
}
