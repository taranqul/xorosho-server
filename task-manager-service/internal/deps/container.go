package deps

import (
	"task-manager-service/internal/domain"
	"task-manager-service/internal/infra/config"

	"go.uber.org/zap"
)

type Container struct {
	service *domain.TaskService
}

func NewContainer(cfg *config.Config, logger *zap.Logger) *Container {
	service, err := domain.NewTaskService()
	if err != nil {
		logger.Sugar().Fatalf("Can't build container %v", err)
	}

	return &Container{
		service: service,
	}

}

func (c *Container) GetTaskService() *domain.TaskService {
	return c.service
}
