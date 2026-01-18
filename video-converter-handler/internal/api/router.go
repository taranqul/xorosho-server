package api

import (
	"video-converter-handler/internal/api/process"
	"video-converter-handler/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	process.NewProcessTaskHandler(logger, eng.Group("/task"), deps.GetConverterService())
	return eng
}
