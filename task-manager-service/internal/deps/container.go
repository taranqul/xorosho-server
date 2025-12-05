package deps

import (
	"context"
	"task-manager-service/internal/domain/service"
	"task-manager-service/internal/infra/config"
	"task-manager-service/internal/infra/repositories"

	"go.uber.org/zap"
)

type Container struct {
	service *service.TaskService
}

func NewContainer(cfg *config.Config, logger *zap.Logger) *Container {
	repository, err := repositories.NewMongoTaskRepository(cfg, context.Background(), logger)
	if err != nil {
		logger.Sugar().Fatalf("Can't build container %v", err)
	}
	service, err := service.NewTaskService(repository)
	if err != nil {
		logger.Sugar().Fatalf("Can't build container %v", err)
	}

	return &Container{
		service: service,
	}

}

func (c *Container) GetTaskService() *service.TaskService {
	return c.service
}
