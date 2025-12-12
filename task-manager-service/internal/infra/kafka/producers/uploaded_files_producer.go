package producers

import (
	"context"
	"encoding/json"
	"task-manager-service/internal/domain"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type TasksProducer struct {
	writer *kafka.Writer
	logger *zap.Logger
}

func NewTasksProducer(addres string, logger *zap.Logger) *TasksProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(addres),
		Topic:    "tasks",
		Balancer: &kafka.LeastBytes{},
	}
	return &TasksProducer{
		writer: writer,
		logger: logger,
	}
}

func (p *TasksProducer) Write(task domain.Task) error {
	jsonBytes, err := json.Marshal(task)
	if err != nil {
		p.logger.Sugar().Errorf("failed to marshal task: %v", err)
		return err
	}

	msg := kafka.Message{
		Key:   []byte(task.Id),
		Value: jsonBytes,
	}

	if err := p.writer.WriteMessages(context.Background(), msg); err != nil {
		p.logger.Sugar().Errorf("failed to write message: %v", err)
		return err
	}
	return nil
}
