package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"stub-nofile-handler/internal/api"
	"stub-nofile-handler/internal/app/server"
	"stub-nofile-handler/internal/deps"
	"stub-nofile-handler/internal/infra/config"

	"go.uber.org/zap"
)

type WorkerRegister struct {
	Name    string         `json:"name"`
	Webhook string         `json:"webhook"`
	Scheme  map[string]any `json:"scheme"`
}

func Run(cfg config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")

	worker := WorkerRegister{
		Name:    "worker1",
		Webhook: "https://example.com/webhook",
		Scheme: map[string]any{
			"task_type": "email",
			"priority":  1,
		},
	}

	data, err := json.Marshal(worker)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://worker-manager-service:8080/worker", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	if resp.Status != fmt.Sprint(200) {
		panic(fmt.Errorf("resposnse code not 200!"))
	}
	container := deps.NewContainer(cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)

}
