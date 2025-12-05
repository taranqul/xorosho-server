package debug

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DebugHandler struct {
	logger *zap.Logger
}

func NewDebugHandler(logger *zap.Logger, rg *gin.RouterGroup) *DebugHandler {
	handler := &DebugHandler{
		logger: logger,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (*DebugHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "I'm healthy!")
}

func (h *DebugHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.GET("/health", h.Health)
}
