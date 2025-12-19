package api

import (
	"worker-manager-service/internal/controllers/api/worker"
	"worker-manager-service/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	worker.NewWorkerHandler(deps.GetService(), logger, eng.Group("/worker"))
	return eng
}
