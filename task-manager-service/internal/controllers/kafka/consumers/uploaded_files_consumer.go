package consumers

import (
	"context"
	"encoding/json"
	"task-manager-service/internal/domain"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type FileService interface {
	HandleUploadedFile(domain.UploadedFilesMessage)
}

type UploadFilesConsumer struct {
	reader  *kafka.Reader
	service FileService
	logger  *zap.Logger
}

func NewUploadFilesConsumer(brokers []string, groupID string, service FileService, logger *zap.Logger) *UploadFilesConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          "uploaded_files",
		MinBytes:       10e3,
		MaxBytes:       10e6,
		CommitInterval: 0,
	})

	return &UploadFilesConsumer{
		reader:  reader,
		service: service,
		logger:  logger,
	}
}

func (c *UploadFilesConsumer) Start(ctx context.Context) {
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
			task, err := decodeUploadedFileMessage(m.Value)

			if err != nil {
				c.logger.Sugar().Error("cant parse message of uploaded files from kafka")
				continue
			}

			go c.service.HandleUploadedFile(task)

			// err = c.reader.CommitMessages(ctx, m)
			// if err != nil {
			// 	c.logger.Sugar().Error("cant commit message of uploaded files from kafka")
			// 	continue
			// }
		}
	}
}

func decodeUploadedFileMessage(value []byte) (domain.UploadedFilesMessage, error) {
	var msg domain.UploadedFilesMessage
	err := json.Unmarshal(value, &msg)
	return msg, err
}
