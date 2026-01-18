package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
	"video-converter-handler/internal/domain/dto"
	"video-converter-handler/internal/infra/repositories"

	"go.uber.org/zap"
)

type ConverterService struct {
	logger          *zap.Logger
	file_repository *repositories.FileRepository
}

func NewConverterService(logger *zap.Logger, file_repository *repositories.FileRepository) *ConverterService {
	return &ConverterService{
		logger:          logger,
		file_repository: file_repository,
	}
}

func (s *ConverterService) ProcessTask(task dto.VideoConverterTask) error {
	s.logger.Sugar().Info(task)

	fileStream, err := s.file_repository.GetFile(task.Objects.Edit)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}
	s.logger.Sugar().Info("Got file:")
	defer fileStream.Close()

	tmpPath := "/tmp/" + task.Objects.Edit
	defer os.Remove(tmpPath)

	err = saveToFile(fileStream, tmpPath)
	if err != nil {
		return fmt.Errorf("failed to save input file: %w", err)
	}

	new_name := fmt.Sprintf("%s_result%s", task.Id, task.Payload.Target)
	tmpResPath := "/tmp/" + new_name

	defer os.Remove(tmpResPath)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	err = convertFile(ctx, tmpPath, tmpResPath)

	if err != nil {
		return fmt.Errorf("failed to convert file: %w", err)
	}

	stream, err := os.Open(tmpResPath)

	if err != nil {
		return fmt.Errorf("failed to get stream file: %w", err)
	}
	defer stream.Close()

	if err := s.file_repository.UploadFile(new_name, stream); err != nil {
		return fmt.Errorf("failed to upload modified file: %w", err)
	}

	s.logger.Sugar().Infof("file uploaded:%s", new_name)

	objects := map[string]string{
		"result": new_name,
	}

	to_send := dto.ResultTask{
		Id:      task.Id,
		Type:    "converter",
		Status:  "completed",
		Objects: objects,
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

func saveToFile(r io.ReadCloser, path string) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

func convertFile(
	ctx context.Context,
	inputPath string,
	outputPath string,
) error {

	cmd := exec.CommandContext(
		ctx,
		"ffmpeg",
		"-y",
		"-i", inputPath,
		outputPath,
	)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	var stderrBuf bytes.Buffer
	go io.Copy(&stderrBuf, stderr)

	err = cmd.Wait()

	if err != nil {
		return fmt.Errorf("ffmpeg failed: %w, stderr: %s", err, stderrBuf.String())
	}

	return nil
}
