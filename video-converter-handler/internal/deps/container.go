package deps

import (
	"video-converter-handler/internal/domain/services"
	"video-converter-handler/internal/infra/config"
	"video-converter-handler/internal/infra/repositories"

	"go.uber.org/zap"
)

type Container struct {
	service *services.ConverterService
}

func NewContainer(cfg config.Config, logger *zap.Logger) *Container {
	file_repository := repositories.NewFileRepository(logger)
	service := services.NewConverterService(logger, file_repository)
	return &Container{
		service: service,
	}

}

func (c *Container) GetConverterService() *services.ConverterService {
	return c.service
}
