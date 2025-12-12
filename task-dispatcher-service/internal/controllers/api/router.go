package api

import (
	"task-dispatcher-service/internal/controllers/api/result"
	"task-dispatcher-service/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	result.NewResultHandler(logger, eng.Group("/task"), deps.GetTaskService())
	return eng
}
