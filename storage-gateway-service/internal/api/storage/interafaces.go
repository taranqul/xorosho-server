package storage

type GatewayService interface {
	GetUploadUrl(filename string, bucketname string) (string, error)
	GetDownloadUrl(filename string, bucketname string) (string, error)
}
