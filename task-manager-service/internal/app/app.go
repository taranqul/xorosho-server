package app

import (
	"fmt"
	"task-manager-service/internal/api"
	"task-manager-service/internal/app/server"
	"task-manager-service/internal/deps"
	"task-manager-service/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	container := deps.NewContainer(cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)
}
