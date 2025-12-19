package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"stub-file-handler/internal/api"
	"stub-file-handler/internal/app/server"
	"stub-file-handler/internal/deps"
	"stub-file-handler/internal/domain/dto"
	"stub-file-handler/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg config.Config, logger *zap.Logger) {

	logger.Info("Trying to establish connection with manager...")
	exe, _ := os.Executable()
	base := filepath.Dir(exe)
	schemaPath := filepath.Join(base, "resources", "schema.json")
	schema_raw, err := os.ReadFile(schemaPath)
	if err != nil {
		panic(err)
	}

	var m map[string]any

	err = json.Unmarshal(schema_raw, &m)
	if err != nil {
		panic(err)
	}

	worker := dto.WorkerRegister{
		Name:    "stub_file",
		Webhook: "http://stub-file-handler:8080/task",
		Scheme:  m,
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
	if resp.StatusCode != 200 {
		panic(fmt.Errorf("resposnse code not 200"))
	}
	logger.Info("Starting wireing...")
	container := deps.NewContainer(cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)

}
