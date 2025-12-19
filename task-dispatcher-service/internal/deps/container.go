package deps

import (
	"context"
	"task-dispatcher-service/internal/domain/services"
	"task-dispatcher-service/internal/infra/config"
	"task-dispatcher-service/internal/infra/kafka/producers"
	"task-dispatcher-service/internal/infra/repositories/webhook"

	"go.uber.org/zap"
)

type Container struct {
	service *services.TaskService
}

func NewContainer(cfg *config.Config, logger *zap.Logger) *Container {
	ctx := context.Background()
	producer := producers.NewResultProducer(cfg.KafkaAddress, logger)
	webhook_repository := webhook.NewRedisRepository(cfg.RedisDSN, ctx, logger)
	service, err := services.NewTaskService(webhook_repository, producer, logger)

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
