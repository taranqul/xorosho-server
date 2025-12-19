package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	LogLevel     string   `env:"LOG_LEVEL" envDefault:"INFO"`
	Port         int      `env:"PORT"`
	GroupID      string   `env:"GROUP_ID"`
	Brokers      []string `env:"BROKERS" envSeparator:","`
	KafkaAddress string   `env:"KAFKA_ADDRESS"`
	RedisDSN     string   `env:"REDIS_DSN" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
