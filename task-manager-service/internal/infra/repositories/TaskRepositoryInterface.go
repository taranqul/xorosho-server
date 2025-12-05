package repositories

import (
	"task-manager-service/internal/domain"

	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	CreateTask(domain.TaskInRepository) (uuid.UUID, error)
}
