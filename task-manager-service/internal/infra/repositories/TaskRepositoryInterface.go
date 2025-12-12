package repositories

import (
	"task-manager-service/internal/domain"

	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	Create(*domain.TaskInRepository) (uuid.UUID, error)
	GetStatus(string) (*string, error)
	Get(string) (*domain.TaskInRepository, error)
	Put(domain.TaskInRepository) error
}
