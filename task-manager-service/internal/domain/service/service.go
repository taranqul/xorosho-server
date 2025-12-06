package service

import (
	"task-manager-service/internal/domain"
	"task-manager-service/internal/infra/repositories"

	"github.com/google/uuid"
)

type TaskService struct {
	task_repository repositories.TaskRepositoryInterface
}

func NewTaskService(task_repository repositories.TaskRepositoryInterface) (*TaskService, error) {
	return &TaskService{
		task_repository: task_repository,
	}, nil
}

func (s *TaskService) CreateTask(task domain.Task) (uuid.UUID, error) {
	to_create := domain.TaskInRepository{
		Id:      task.Id.String(),
		Type:    task.Type,
		Objects: task.Objects,
		Payload: task.Payload,
		Status:  "idle",
	}
	return s.task_repository.CreateTask(to_create)
}

func (s *TaskService) GetTaskStatus(id string) (string, error) {
	return s.task_repository.GetTaskStatus(id)
}
