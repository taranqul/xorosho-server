package consumers

import (
	"context"
	"encoding/json"
	"task-manager-service/internal/domain"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type TaskResultService interface {
	HandleResult(domain.TaskInRepository)
}

type TaskResultConsumer struct {
	reader  *kafka.Reader
	service TaskResultService
	logger  *zap.Logger
}

func NewTaskResultConsumer(brokers []string, groupID string, service TaskResultService, logger *zap.Logger) *TaskResultConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          "task_result",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: 0,
	})

	return &TaskResultConsumer{
		reader:  reader,
		service: service,
		logger:  logger,
	}
}

func (c *TaskResultConsumer) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				c.logger.Sugar().Error("cant read message of uploaded files from kafka")
				continue
			}
			task, err := decodeTaskMessage(m.Value)

			if err != nil {
				c.logger.Sugar().Error("cant parse message of uploaded files from kafka")
				continue
			}

			go c.service.HandleResult(task)

			// err = c.reader.CommitMessages(ctx, m)
			// if err != nil {
			// 	c.logger.Sugar().Error("cant commit message of uploaded files from kafka")
			// 	continue
			// }
		}
	}
}

func decodeTaskMessage(value []byte) (domain.TaskInRepository, error) {
	var msg domain.TaskInRepository
	err := json.Unmarshal(value, &msg)
	return msg, err
}
