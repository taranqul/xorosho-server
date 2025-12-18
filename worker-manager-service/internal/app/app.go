package app

import (
	"fmt"
	"worker-manager-service/internal/app/server"
	"worker-manager-service/internal/controllers/api"
	"worker-manager-service/internal/infra/config"

	"worker-manager-service/internal/deps"

	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	container := deps.NewContainer(*cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)
}
