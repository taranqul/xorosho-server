package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	MongoURI string `env:"MONGO_URI" required:"true"`
	MongoDB  string `env:"MONGO_DB" required:"true"`
	RedisDSN string `envconfig:"REDIS_DSN" required:"true"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
	Port     int    `env:"PORT"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
