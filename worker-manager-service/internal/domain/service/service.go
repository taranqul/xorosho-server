package service

import (
	"worker-manager-service/internal/domain/dto"
	"worker-manager-service/internal/infra/repositories/webhook"
	"worker-manager-service/internal/infra/repositories/worker"

	"github.com/xeipuuv/gojsonschema"
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
	err := w.webhook_repository.Set(worker.Name, worker.Webhook)
	if err != nil {
		return err
	}

	err = w.worker_repository.Create(&worker)
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkerService) ValidateScheme(request dto.TaskRequest) (bool, error) {
	merged := map[string]any{
		"objects": request.Objects,
		"payload": request.Payload,
	}
	worker, err := w.worker_repository.Read(request.Type)
	if err != nil {
		return false, err
	}

	if worker == nil {
		return false, nil
	}

	schemaLoader := gojsonschema.NewGoLoader(worker.Scheme)
	documentLoader := gojsonschema.NewGoLoader(merged)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return false, err
	}

	if !result.Valid() {
		return false, nil
	}

	return true, nil
}

func (w *WorkerService) GetTasks() ([]string, error) {
	return w.webhook_repository.Scan()
}

func (w *WorkerService) GetScheme(name string) (map[string]any, error) {
	worker, err := w.worker_repository.Read(name)

	if err != nil || worker == nil {
		return nil, err
	}

	return worker.Scheme, nil
}
