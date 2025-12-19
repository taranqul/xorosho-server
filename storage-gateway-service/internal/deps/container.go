package deps

import (
	"context"
	"storage-gateway-service/internal/api/storage"
	"storage-gateway-service/internal/domain"
	"storage-gateway-service/internal/infra/config"
	"storage-gateway-service/internal/infra/minio"
)

type Container struct {
	service storage.GatewayService
}

func NewContainer(cfg config.Config) *Container {
	ctx := context.Background()
	ext_minio_storage := minio.NewMinioDAO(cfg, cfg.MinioExternalEndpoint, cfg.MinioEndpoint, &ctx)
	int_minio_storage := minio.NewMinioDAO(cfg, cfg.MinioEndpoint, cfg.MinioEndpoint, &ctx)
	service := domain.NewGatewayService(ext_minio_storage, int_minio_storage)
	return &Container{
		service: service,
	}

}

func (c *Container) GetGatewayService() storage.GatewayService {
	return c.service
}
