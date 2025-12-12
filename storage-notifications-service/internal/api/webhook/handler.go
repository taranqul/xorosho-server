package webhook

import (
	"net/http"
	"storage-notifications-service/internal/domain/dto"
	"storage-notifications-service/internal/domain/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebhookHandler struct {
	logger  *zap.Logger
	service *services.UploadedFilesService
}

func NewWebhookHandler(service *services.UploadedFilesService, logger *zap.Logger, rg *gin.RouterGroup) *WebhookHandler {
	handler := &WebhookHandler{
		logger:  logger,
		service: service,
	}
	handler.registerEndpoints(rg)
	return handler
}

func (h *WebhookHandler) Webhook(c *gin.Context) {
	var event dto.S3Event

	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(event.Records) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no records found"})
		return
	}

	objectKey := event.Records[0].S3.Object.Key
	err := h.service.HandleUploadedFile(objectKey)
	if err != nil {
		h.logger.Sugar().Errorf("Something goes wrong while trying to process object key: %v", err)
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *WebhookHandler) registerEndpoints(rg *gin.RouterGroup) {
	rg.POST("/", h.Webhook)
}
