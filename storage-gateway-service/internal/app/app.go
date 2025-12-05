package app

import (
	"fmt"
	"storage-gateway-service/internal/api"
	"storage-gateway-service/internal/app/server"
	"storage-gateway-service/internal/deps"
	"storage-gateway-service/internal/infra/config"

	"go.uber.org/zap"
)

func Run(cfg config.Config, logger *zap.Logger) {
	logger.Info("Starting wireing...")
	logger.Info("Endpoint:" + cfg.MinioEndpoint)
	container := deps.NewContainer(cfg)
	eng := api.RegisterHandlers(logger, container)
	server.StartServer(eng, fmt.Sprint(cfg.Port), logger)

}
