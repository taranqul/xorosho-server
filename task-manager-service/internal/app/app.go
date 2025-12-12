package app

import (
	"fmt"
	"task-manager-service/internal/app/server"
	"task-manager-service/internal/controllers/api"
	"task-manager-service/internal/controllers/kafka"
	"task-manager-service/internal/deps"
	"task-manager-service/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	container := deps.NewContainer(cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	kafka_orchestrator := kafka.NewOrchestrator(cfg, logger, container)
	kafka_orchestrator.Start()
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)
}
