package producers

import (
	"context"
	"encoding/json"
	"storage-notifications-service/internal/domain/dto"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type UploadedFilesProducer struct {
	writer *kafka.Writer
	logger *zap.Logger
}

func NewUploadedFilesProducer(addres string, logger *zap.Logger) *UploadedFilesProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(addres),
		Topic:    "uploaded_files",
		Balancer: &kafka.LeastBytes{},
	}
	return &UploadedFilesProducer{
		writer: writer,
		logger: logger,
	}
}

func (p *UploadedFilesProducer) Write(task dto.UploadedFilesMessage) error {
	jsonBytes, err := json.Marshal(task)
	if err != nil {
		p.logger.Sugar().Errorf("failed to marshal task: %v", err)
		return err
	}

	msg := kafka.Message{
		Key:   []byte(task.TaskID),
		Value: jsonBytes,
	}

	if err := p.writer.WriteMessages(context.Background(), msg); err != nil {
		p.logger.Sugar().Errorf("failed to write message: %v", err)
		return err
	}
	return nil
}
