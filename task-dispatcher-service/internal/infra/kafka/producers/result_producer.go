package producers

import (
	"context"
	"encoding/json"
	"task-dispatcher-service/internal/domain/dto"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type ResultProducer struct {
	writer *kafka.Writer
	logger *zap.Logger
}

func NewResultProducer(addres string, logger *zap.Logger) *ResultProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(addres),
		Topic:    "task_result",
		Balancer: &kafka.LeastBytes{},
	}
	return &ResultProducer{
		writer: writer,
		logger: logger,
	}
}

func (p *ResultProducer) Write(task dto.Task) error {
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
