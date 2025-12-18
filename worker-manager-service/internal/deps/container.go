package deps

import (
	"context"
	"worker-manager-service/internal/domain/service"
	"worker-manager-service/internal/infra/config"
	"worker-manager-service/internal/infra/repositories/webhook"
	"worker-manager-service/internal/infra/repositories/worker"

	"go.uber.org/zap"
)

type Container struct {
	worker_service *service.WorkerService
}

func NewContainer(cfg config.Config, logger *zap.Logger) *Container {
	ctx := context.Background()
	worker_repo, err := worker.NewMongoRepository(ctx, logger, cfg.MongoURI, cfg.MongoDB)
	if err != nil {
		panic(err)
	}
	webhook_repo, err := webhook.NewRedisRepository(cfg.RedisDSN, ctx, logger)
	if err != nil {
		panic(err)
	}
	worker_service, err := service.NewWorkerService(webhook_repo, worker_repo, logger)
	return &Container{worker_service: worker_service}

}

func (c *Container) GetService() *service.WorkerService {
	return c.worker_service
}
