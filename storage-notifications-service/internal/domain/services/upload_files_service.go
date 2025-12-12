package services

import (
	"errors"
	"storage-notifications-service/internal/domain/dto"
	"storage-notifications-service/internal/infra/kafka/producers"
	"strings"

	"go.uber.org/zap"
)

type UploadedFilesService struct {
	producer *producers.UploadedFilesProducer
	logger   *zap.Logger
}

func NewUploadedFilesService(producer *producers.UploadedFilesProducer, logger *zap.Logger) *UploadedFilesService {
	return &UploadedFilesService{
		producer: producer,
		logger:   logger,
	}
}

func (s *UploadedFilesService) HandleUploadedFile(file string) error {
	id, typ, err := parseFilename(file)
	if err != nil {
		return err
	}

	task := &dto.UploadedFilesMessage{
		TaskID: id,
		Type:   typ,
		File:   file,
	}

	err = s.producer.Write(*task)

	if err != nil {
		return err
	}

	return nil
}

func parseFilename(filename string) (id, typ string, err error) {
	dotIdx := strings.LastIndex(filename, ".")
	if dotIdx == -1 {
		err = errors.New("ext missing")
		return
	}
	filename = filename[:dotIdx]

	parts := strings.SplitN(filename, "_", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		err = errors.New("id or typ missing")
		return
	}

	id = parts[0]
	typ = parts[1]
	return
}
