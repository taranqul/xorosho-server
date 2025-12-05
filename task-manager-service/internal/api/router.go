package api

import (
	"task-manager-service/internal/api/debug"
	"task-manager-service/internal/api/task"
	"task-manager-service/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	debug.NewDebugHandler(logger, eng.Group("/debug"))
	task.NewTaskHandler(deps.GetTaskService(), logger, eng.Group("/task"))
	return eng
}
