package services

import (
	"task-manager-service/internal/domain"
	"task-manager-service/internal/infra/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type TaskService struct {
	task_repository repositories.TaskRepositoryInterface
	logger          *zap.Logger
}

func NewTaskService(task_repository repositories.TaskRepositoryInterface, logger *zap.Logger) (*TaskService, error) {
	return &TaskService{
		task_repository: task_repository,
		logger:          logger,
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
	return s.task_repository.Create(&to_create)
}

func (s *TaskService) GetTaskStatus(id string) (*string, error) {
	return s.task_repository.GetStatus(id)
}

func (s *TaskService) HandleUploadedFile(uploaded_file domain.UploadedFilesMessage) {
	task, err := s.task_repository.Get(uploaded_file.TaskID)
	if err != nil {
		s.logger.Sugar().Errorf("uploaded file with %v wasnt getted because of %v, but was commited (thats actualy very bad)", uploaded_file.File, err)
		return
	}

	task.Objects[uploaded_file.Type] = uploaded_file.File
	err = s.task_repository.Put(*task)
	if err != nil {
		s.logger.Sugar().Errorf("uploaded file with %v wasnt putted because of %v, but was commited (thats actualy very bad)", uploaded_file.File, err)
		return
	}
}
