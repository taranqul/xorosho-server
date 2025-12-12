package app

import (
	"fmt"
	"stub-nofile-handler/internal/api"
	"stub-nofile-handler/internal/app/server"
	"stub-nofile-handler/internal/deps"
	"stub-nofile-handler/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	container := deps.NewContainer(cfg, logger)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)

}
