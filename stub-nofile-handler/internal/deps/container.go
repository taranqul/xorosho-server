package deps

import (
	"stub-nofile-handler/internal/domain/services"
	"stub-nofile-handler/internal/infra/config"

	"go.uber.org/zap"
)

type Container struct {
	service *services.StubService
}

func NewContainer(cfg config.Config, logger *zap.Logger) *Container {
	service := services.NewStubService(logger)
	return &Container{
		service: service,
	}

}

func (c *Container) GetStubService() *services.StubService {
	return c.service
}
