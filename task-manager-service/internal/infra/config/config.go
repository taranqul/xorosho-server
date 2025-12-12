package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	LogLevel string   `env:"LOG_LEVEL" envDefault:"INFO"`
	MongoURI string   `env:"MONGO_URI"`
	MongoDB  string   `env:"MONGO_DB"`
	Port     int      `env:"PORT"`
	GroupID  string   `env:"GROUP_ID"`
	Brokers  []string `env:"BROKERS" envSeparator:","`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
