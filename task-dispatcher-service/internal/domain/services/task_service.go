package services

import (
	"task-dispatcher-service/internal/domain/dto"

	"go.uber.org/zap"
)

type TaskService struct {
	logger *zap.Logger
}

func NewTaskService(logger *zap.Logger) (*TaskService, error) {
	return &TaskService{
		logger: logger,
	}, nil
}

func (s *TaskService) DispatchTask(task dto.Task) {
	s.logger.Info(task.Id)
}
