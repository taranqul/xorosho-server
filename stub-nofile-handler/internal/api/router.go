package api

import (
	"stub-nofile-handler/internal/api/process"
	"stub-nofile-handler/internal/deps"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterHandlers(logger *zap.Logger, deps *deps.Container) *gin.Engine {
	eng := gin.Default()
	process.NewProcessTaskHandler(logger, eng.Group("/task"), deps.GetStubService())
	return eng
}
