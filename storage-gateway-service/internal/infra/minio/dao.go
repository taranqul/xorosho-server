package minio

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"storage-gateway-service/internal/infra/config"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioDAO struct {
	client *minio.Client
	ctx    *context.Context
}

func NewMinioDAO(cfg config.Config, MinioExternalEndpoint string, MinioEndpoint string, ctx *context.Context) *MinioDAO {
	tr := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			if strings.Contains(addr, MinioExternalEndpoint) {
				addr = MinioEndpoint
			}
			return (&net.Dialer{}).DialContext(ctx, "tcp4", addr)
		},
	}

	client, err := minio.New(MinioExternalEndpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretAccessKey, ""),
		Secure:    cfg.MinIoUseSSL,
		Region:    "us-east-1",
		Transport: tr,
	})
	if err != nil {
		panic(err)
	}
	return &MinioDAO{client, ctx}
}

func (m *MinioDAO) GetBuckets() error {
	buckets, err := m.client.ListBuckets(*m.ctx)
	if err != nil {
		return err
	}
	fmt.Printf("buckets: %v\n", buckets)
	return nil
}

func (m *MinioDAO) GetUploadUrl(filename string, bucketname string) (string, error) {
	url, err := m.client.PresignedPutObject(*m.ctx, bucketname, filename, time.Minute*10)
	if err != nil {
		print(err)
		return "", err
	}
	return url.String(), nil
}

func (m *MinioDAO) GetDownloadUrl(filename string, bucketname string) (string, error) {
	url, err := m.client.PresignedGetObject(*m.ctx, bucketname, filename, time.Minute*10, nil)
	if err != nil {
		print(err)
		return "", err
	}
	return url.String(), nil
}
