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

func (h *StorageHandler) GetExtUploadUrl(c *gin.Context) {
	filename := c.Query("filename")
	bucketname := c.Query("bucketname")
	url, err := h.service.GetExtUploadUrl(filename, bucketname)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *StorageHandler) GetExtDownloadUrl(c *gin.Context) {
	filename := c.Query("filename")
	bucketname := c.Query("bucketname")
	url, err := h.service.GetExtDownloadUrl(filename, bucketname)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *StorageHandler) GetIntUploadUrl(c *gin.Context) {
	filename := c.Query("filename")
	bucketname := c.Query("bucketname")
	url, err := h.service.GetIntUploadUrl(filename, bucketname)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *StorageHandler) GetIntDownloadUrl(c *gin.Context) {
	filename := c.Query("filename")
	bucketname := c.Query("bucketname")
	url, err := h.service.GetIntDownloadUrl(filename, bucketname)
	if err != nil {
		c.JSON(501, err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *StorageHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.GET("/external/uploadUrl", h.GetExtUploadUrl)
	rg.GET("/external/downloadUrl", h.GetExtDownloadUrl)
	rg.GET("/internal/uploadUrl", h.GetIntUploadUrl)
	rg.GET("/internal/downloadUrl", h.GetIntDownloadUrl)
}
