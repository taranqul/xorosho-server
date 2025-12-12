package kafka

import (
	"context"
	"sync"
	"task-dispatcher-service/internal/controllers/kafka/consumers"
	"task-dispatcher-service/internal/deps"
	"task-dispatcher-service/internal/infra/config"

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
	upload_consumer := consumers.NewInitiatedTaskConsumer(cfg.Brokers, cfg.GroupID, deps.GetTaskService(), logger)
	cons = append(cons, upload_consumer)
	return &KafkaOrchestrator{
		ctx:       ctx,
		consumers: cons,
	}
}

func (o *KafkaOrchestrator) Start() {
	var wg sync.WaitGroup

	for _, c := range o.consumers {
		wg.Add(1)
		go func(c Consumer) {
			defer wg.Done()
			c.Start(o.ctx)
		}(c)
	}

	wg.Wait()
}
