package api

import (
	"storage-gateway-service/internal/api/debug"

	"storage-gateway-service/internal/api/storage"
	"storage-gateway-service/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	debug.NewDebugHandler(logger, eng.Group("/debug"))
	storage.NewStorageHandler(logger, deps.GetGatewayService(), eng.Group("/storage"))
	return eng
}
