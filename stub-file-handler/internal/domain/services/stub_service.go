package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stub-file-handler/internal/domain/dto"
	"stub-file-handler/internal/infra/repositories"

	"go.uber.org/zap"
)

type StubService struct {
	logger          *zap.Logger
	file_repository *repositories.FileRepository
}

func NewStubService(logger *zap.Logger, file_repository *repositories.FileRepository) *StubService {
	return &StubService{
		logger:          logger,
		file_repository: file_repository,
	}
}

func (s *StubService) ProcessTask(task dto.StubFileTask) error {
	s.logger.Sugar().Info(task)

	fileStream, err := s.file_repository.GetFile(task.Objects.Edit)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}
	s.logger.Sugar().Info("Got file:")
	defer fileStream.Close()

	modifiedStream, err := stubModify(fileStream, task.Payload.Trace)

	if err != nil {
		return fmt.Errorf("failed to modify file: %w", err)
	}
	new_name := fmt.Sprint(task.Id, "_result.txt")

	if err := s.file_repository.UploadFile(new_name, modifiedStream); err != nil {
		return fmt.Errorf("failed to upload modified file: %w", err)
	}

	s.logger.Sugar().Infof("file uploaded:%s", new_name)
	res := dto.Result{
		Trace: fmt.Sprintf("hello from stub, trace: %s. File: %s", task.Payload.Trace, task.Objects.Edit),
	}
	objects := map[string]string{
		"result": new_name,
	}
	to_send := dto.ResultTask{
		Id:      task.Id,
		Type:    "stub_no_file",
		Status:  "completed",
		Objects: objects,
		Payload: res,
	}

	jsonData, err := json.Marshal(to_send)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://task-dispatcher-service:8080/task", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s.logger.Sugar().Info(resp.Status)
	return nil
}

func stubModify(r io.Reader, to_write string) (io.Reader, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	data = append(data, []byte("\nModified by service!")...)
	data = append(data, []byte("\n"+to_write)...)
	return bytes.NewReader(data), nil
}
