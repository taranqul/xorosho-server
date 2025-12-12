package main

import (
	"stub-nofile-handler/internal/app"
	"stub-nofile-handler/internal/infra/config"

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
	app.Run(*cfg, logger)
}
