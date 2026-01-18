package main

import (
	"video-converter-handler/internal/app"
	"video-converter-handler/internal/infra/config"

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
