package main

import (
	"task-manager-service/internal/app"
	"task-manager-service/internal/infra/config"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	app.Run(cfg, logger)
}
