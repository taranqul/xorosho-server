package task

import (
	"task-manager-service/internal/domain"
	"task-manager-service/internal/domain/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskHandler struct {
	logger  *zap.Logger
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService, logger *zap.Logger, rg *gin.RouterGroup) *TaskHandler {
	handler := &TaskHandler{
		logger:  logger,
		service: service,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task domain.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uuid, err := h.service.CreateTask(task)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"uuid": uuid})
}

func (h *TaskHandler) GetTaskStatus(c *gin.Context) {
	id := c.Query("id")
	status, err := h.service.GetTaskStatus(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"status": status})
}

func (h *TaskHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("", h.CreateTask)
	rg.GET("/status", h.GetTaskStatus)
}
