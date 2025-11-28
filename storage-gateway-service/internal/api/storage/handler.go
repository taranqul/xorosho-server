package storage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StorageHandler struct {
	logger  *zap.Logger
	service GatewayService
}

func NewStorageHandler(logger *zap.Logger, service GatewayService, rg *gin.RouterGroup) *StorageHandler {
	handler := &StorageHandler{
		logger:  logger,
		service: service,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (h *StorageHandler) GetUploadUrl(c *gin.Context) {
	filename := c.Query("filename")
	bucketname := c.Query("bucketname")
	url, err := h.service.GetUploadUrl(filename, bucketname)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *StorageHandler) GetDownloadUrl(c *gin.Context) {
	filename := c.Query("filename")
	bucketname := c.Query("bucketname")
	url, err := h.service.GetDownloadUrl(filename, bucketname)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *StorageHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.GET("/uploadUrl", h.GetUploadUrl)
	rg.GET("/downloadUrl", h.GetDownloadUrl)
}
