package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"task-dispatcher-service/internal/domain/dto"
	"task-dispatcher-service/internal/infra/kafka/producers"

	"go.uber.org/zap"
)

type TaskService struct {
	logger          *zap.Logger
	result_producer *producers.ResultProducer
}

func NewTaskService(result_producer *producers.ResultProducer, logger *zap.Logger) (*TaskService, error) {
	return &TaskService{
		logger:          logger,
		result_producer: result_producer,
	}, nil
}

func (s *TaskService) DispatchTask(task dto.Task) {

	jsonData, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://stub-nofile-handler:8080/task", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s.logger.Sugar().Info(resp.Status)
}

func (s *TaskService) HandleResult(task dto.Task) {
	s.result_producer.Write(task)
}
