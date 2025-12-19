package worker

import "worker-manager-service/internal/domain/dto"

type IWorkerRepository interface {
	Create(*dto.WorkerRegister) error
	Read(string) (*dto.WorkerRegister, error)
}
