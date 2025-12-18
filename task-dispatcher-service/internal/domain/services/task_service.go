package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"task-dispatcher-service/internal/domain/dto"
	"task-dispatcher-service/internal/infra/kafka/producers"
	"task-dispatcher-service/internal/infra/repositories/webhook"

	"go.uber.org/zap"
)

type TaskService struct {
	logger             *zap.Logger
	result_producer    *producers.ResultProducer
	webhook_repository webhook.IWebhookRepository
}

func NewTaskService(webhook_repository webhook.IWebhookRepository, result_producer *producers.ResultProducer, logger *zap.Logger) (*TaskService, error) {
	return &TaskService{
		logger:             logger,
		result_producer:    result_producer,
		webhook_repository: webhook_repository,
	}, nil
}

func (s *TaskService) DispatchTask(task dto.Task) {

	jsonData, err := json.Marshal(task)
	if err != nil {
		return
	}

	webhook, err := s.webhook_repository.Get(task.Type)
	if err != nil || webhook == "" {
		return
	}

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	s.logger.Sugar().Info(resp.Status)
}

func (s *TaskService) HandleResult(task dto.Task) {
	s.result_producer.Write(task)
}
