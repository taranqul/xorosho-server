package domain

import "github.com/google/uuid"

type TaskService struct {
}

func NewTaskService() (*TaskService, error) {
	return &TaskService{}, nil
}

func (*TaskService) CreateTask(task Task) (uuid.UUID, error) {
	return task.Id, nil
}
