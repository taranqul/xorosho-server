package api

import (
	"storage-notifications-service/internal/api/debug"
	"storage-notifications-service/internal/api/webhook"
	"storage-notifications-service/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	debug.NewDebugHandler(logger, eng.Group("/debug"))
	webhook.NewWebhookHandler(deps.GetUploadedFilesService(), logger, eng.Group("/webhook"))
	return eng
}
