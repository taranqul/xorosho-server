package deps

import (
	"stub-file-handler/internal/domain/services"
	"stub-file-handler/internal/infra/config"
	"stub-file-handler/internal/infra/repositories"

	"go.uber.org/zap"
)

type Container struct {
	service *services.StubService
}

func NewContainer(cfg config.Config, logger *zap.Logger) *Container {
	file_repository := repositories.NewFileRepository(logger)
	service := services.NewStubService(logger, file_repository)
	return &Container{
		service: service,
	}

}

func (c *Container) GetStubService() *services.StubService {
	return c.service
}
