package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"stub-nofile-handler/internal/domain/dto"

	"go.uber.org/zap"
)

type StubService struct {
	logger *zap.Logger
}

func NewStubService(logger *zap.Logger) *StubService {
	return &StubService{
		logger: logger,
	}
}

func (s *StubService) ProcessTask(task dto.StubNoFileTask) error {
	res := dto.Result{
		Trace: "hello from stub ",
	}

	to_send := dto.ResultTask{
		Id:      task.Id,
		Type:    "stub_no_file",
		Status:  "completed",
		Objects: map[string]string{},
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
