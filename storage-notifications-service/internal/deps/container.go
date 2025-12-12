package deps

import (
	"storage-notifications-service/internal/domain/services"
	"storage-notifications-service/internal/infra/config"
	"storage-notifications-service/internal/infra/kafka/producers"

	"go.uber.org/zap"
)

type Container struct {
	upload_files_service *services.UploadedFilesService
}

func NewContainer(cfg config.Config, logger *zap.Logger) *Container {
	upload_files_producer := producers.NewUploadedFilesProducer(cfg.KafkaAddress, logger)
	upload_files_service := services.NewUploadedFilesService(
		upload_files_producer,
		logger,
	)
	return &Container{
		upload_files_service: upload_files_service,
	}

}

func (c *Container) GetUploadedFilesService() *services.UploadedFilesService {
	return c.upload_files_service
}
