package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	LogLevel     string `env:"LOG_LEVEL" envDefault:"INFO"`
	KafkaAddress string `env:"KAFKA_ADDRESS"`
	Port         int    `env:"PORT"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
