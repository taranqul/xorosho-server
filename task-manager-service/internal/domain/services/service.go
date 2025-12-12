package services

import (
	"task-manager-service/internal/domain"
	"task-manager-service/internal/infra/kafka/producers"
	"task-manager-service/internal/infra/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type TaskService struct {
	task_repository repositories.TaskRepositoryInterface
	task_producer   *producers.TasksProducer
	logger          *zap.Logger
}

func NewTaskService(task_producer *producers.TasksProducer, task_repository repositories.TaskRepositoryInterface, logger *zap.Logger) (*TaskService, error) {
	return &TaskService{
		task_producer:   task_producer,
		task_repository: task_repository,
		logger:          logger,
	}, nil
}

func (s *TaskService) CreateTask(task domain.Task) (*uuid.UUID, error) {
	ready_to_init := false
	status := "idle"

	for i := range task.Objects {
		task.Objects[i] = ""
	}

	if len(task.Objects) == 0 {
		status = "in progress"
		ready_to_init = true
	}

	to_create := domain.TaskInRepository{
		Id:      task.Id,
		Type:    task.Type,
		Objects: task.Objects,
		Payload: task.Payload,
		Status:  status,
	}
	uuid, err := s.task_repository.Create(&to_create)

	if err != nil {

		s.logger.Sugar().Errorf("task wasnt created: %v", to_create.Id)
		return nil, err
	}

	if ready_to_init {
		s.initTask(&to_create)
	}

	return &uuid, nil
}

func (s *TaskService) GetTaskStatus(id string) (*string, error) {
	return s.task_repository.GetStatus(id)
}

func (s *TaskService) GetTask(id string) (*domain.TaskInRepository, error) {
	return s.task_repository.Get(id)
}

func (s *TaskService) HandleUploadedFile(uploaded_file domain.UploadedFilesMessage) {
	task, err := s.task_repository.Get(uploaded_file.TaskID)
	if err != nil {
		s.logger.Sugar().Errorf("uploaded file with %v wasnt getted because of %v, but was commited (thats actualy very bad)", uploaded_file.File, err)
		return
	}

	task.Objects[uploaded_file.Type] = uploaded_file.File
	ready_to_init := true

	for _, v := range task.Objects {
		if v == "" {
			ready_to_init = false
			break
		}
	}
	if ready_to_init {
		task.Status = "In progress"
	}

	err = s.task_repository.Put(*task)
	if err != nil {
		s.logger.Sugar().Errorf("uploaded file with %v wasnt putted because of %v, but was commited (thats actualy very bad) and marked as 'in progres' (even worse)", uploaded_file.File, err)
		return
	}

	if ready_to_init {
		s.initTask(task)
	}
}

func (s *TaskService) HandleResult(task domain.TaskInRepository) {
	err := s.task_repository.Put(task)
	if err != nil {
		s.logger.Sugar().Errorf("result with %v cant putted: %v", task.Id, err)
		return
	}
}

func (s *TaskService) initTask(task *domain.TaskInRepository) {
	s.logger.Sugar().Infof("not implemented yet: %v", task.Id)
	err := s.task_producer.Write(domain.Task{
		Id:      task.Id,
		Type:    task.Type,
		Objects: task.Objects,
		Payload: task.Payload,
	})

	if err != nil {
		s.logger.Sugar().Errorf("task with %v wasnt initiated", task.Id)
		return
	}
}
