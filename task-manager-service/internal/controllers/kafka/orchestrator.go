package kafka

import (
	"context"
	"task-manager-service/internal/controllers/kafka/consumers"
	"task-manager-service/internal/deps"
	"task-manager-service/internal/infra/config"

	"go.uber.org/zap"
)

type Consumer interface {
	Start(context.Context)
}

type KafkaOrchestrator struct {
	ctx       context.Context
	consumers []Consumer
}

func NewOrchestrator(cfg *config.Config, logger *zap.Logger, deps *deps.Container) *KafkaOrchestrator {
	ctx := context.Background()
	var cons []Consumer
	upload_consumer := consumers.NewUploadFilesConsumer(cfg.Brokers, cfg.GroupID, deps.GetTaskService(), logger)
	result_consumer := consumers.NewTaskResultConsumer(cfg.Brokers, cfg.GroupID, deps.GetTaskService(), logger)
	cons = append(cons, upload_consumer, result_consumer)
	return &KafkaOrchestrator{
		ctx:       ctx,
		consumers: cons,
	}
}

func (o *KafkaOrchestrator) Start() {
	for _, c := range o.consumers {
		go func(c Consumer) {
			c.Start(o.ctx)
		}(c)
	}
}
