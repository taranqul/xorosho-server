package deps

import (
	"task-dispatcher-service/internal/domain/services"
	"task-dispatcher-service/internal/infra/config"

	"go.uber.org/zap"
)

type Container struct {
	service *services.TaskService
}

func NewContainer(cfg *config.Config, logger *zap.Logger) *Container {

	service, err := services.NewTaskService(logger)

	if err != nil {
		logger.Sugar().Fatalf("Can't build container %v", err)
	}

	return &Container{
		service: service,
	}

}

func (c *Container) GetTaskService() *services.TaskService {
	return c.service
}
