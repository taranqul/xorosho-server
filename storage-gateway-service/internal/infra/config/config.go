package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	MinioEndpoint         string `env:"MINIO_ENDOINT,required"`
	MinioExternalEndpoint string `env:"MINIO_EXTERNAL_ENDOINT,required"`
	MinioAccessKeyID      string `env:"MINIO_ACCESS_KEY_ID,required"`
	MinioSecretAccessKey  string `env:"MINIO_SECRET_KEY,required"`
	MinIoUseSSL           bool   `env:"MINIO_USE_SSL" envDefault:"true"`
	LogLevel              string `env:"LOG_LEVEL" envDefault:"INFO"`
	Port                  int    `env:"PORT"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
