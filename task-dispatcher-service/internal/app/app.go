package app

import (
	"task-dispatcher-service/internal/controllers/kafka"
	"task-dispatcher-service/internal/deps"
	"task-dispatcher-service/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	container := deps.NewContainer(cfg, logger)
	kafka_orchestrator := kafka.NewOrchestrator(cfg, logger, container)
	kafka_orchestrator.Start()
}
