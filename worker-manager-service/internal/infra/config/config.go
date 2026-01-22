package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	MongoURI       string `env:"MONGO_URI" required:"true"`
	MongoDB        string `env:"MONGO_DB" required:"true"`
	RedisDSN       string `env:"REDIS_DSN" required:"true"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"INFO"`
	WorkerLifeTime int    `env:"WORKER_LIFE_TIME" envDefault:"30"`
	Port           int    `env:"PORT"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
