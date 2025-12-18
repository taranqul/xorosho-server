package service

import (
	"worker-manager-service/internal/domain/dto"
	"worker-manager-service/internal/infra/repositories/webhook"
	"worker-manager-service/internal/infra/repositories/worker"

	"go.uber.org/zap"
)

type WorkerService struct {
	webhook_repository webhook.IWebhookRepository
	worker_repository  worker.IWorkerRepository
	logger             *zap.Logger
}

func NewWorkerService(webhook_repository webhook.IWebhookRepository, worker_repository worker.IWorkerRepository, logger *zap.Logger) (*WorkerService, error) {
	return &WorkerService{
		webhook_repository: webhook_repository,
		worker_repository:  worker_repository,
		logger:             logger,
	}, nil
}

func (w *WorkerService) RegisterService(worker dto.WorkerRegister) error {
	return w.webhook_repository.Set(worker.Name, worker.Webhook)
}
