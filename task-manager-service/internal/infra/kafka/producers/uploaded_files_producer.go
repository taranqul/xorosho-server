package producers

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"task-manager-service/internal/domain"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type TasksProducer struct {
	writer *kafka.Writer
	logger *zap.Logger
}

func NewTasksProducer(addres string, logger *zap.Logger) *TasksProducer {
	writer := &kafka.Writer{
		Addr:        kafka.TCP(addres),
		Topic:       "tasks",
		Balancer:    &kafka.LeastBytes{},
		Logger:      log.New(os.Stdout, "kafka writer: ", log.LstdFlags),
		ErrorLogger: log.New(os.Stderr, "kafka error: ", log.LstdFlags),
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		p.logger.Sugar().Errorf("failed to write message: %v", err)
		return err
	}
	return nil
}
