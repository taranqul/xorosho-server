package consumers

import (
	"context"
	"encoding/json"
	"task-dispatcher-service/internal/domain/dto"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type TaskService interface {
	DispatchTask(dto.Task)
}

type InitiatedTaskConsumer struct {
	reader  *kafka.Reader
	service TaskService
	logger  *zap.Logger
}

func NewInitiatedTaskConsumer(brokers []string, groupID string, service TaskService, logger *zap.Logger) *InitiatedTaskConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          "tasks",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: 0,
	})

	return &InitiatedTaskConsumer{
		reader:  reader,
		service: service,
		logger:  logger,
	}
}

func (c *InitiatedTaskConsumer) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				c.logger.Sugar().Error("cant read message of initiated task from kafka")
				continue
			}
			task, err := decodeUploadedFileMessage(m.Value)

			if err != nil {
				c.logger.Sugar().Error("cant parse message of uploaded files from kafka")
				continue
			}

			go c.service.DispatchTask(task)

			// err = c.reader.CommitMessages(ctx, m)
			// if err != nil {
			// 	c.logger.Sugar().Error("cant commit message of uploaded files from kafka")
			// 	continue
			// }
		}
	}
}

func decodeUploadedFileMessage(value []byte) (dto.Task, error) {
	var msg dto.Task
	err := json.Unmarshal(value, &msg)
	return msg, err
}
