package app

import (
	"fmt"
	"storage-notifications-service/internal/api"
	"storage-notifications-service/internal/app/server"
	"storage-notifications-service/internal/deps"
	"storage-notifications-service/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	container := deps.NewContainer(cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)

}
