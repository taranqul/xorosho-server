package storage

type GatewayService interface {
	GetExtUploadUrl(filename string, bucketname string) (string, error)
	GetExtDownloadUrl(filename string, bucketname string) (string, error)
	GetIntUploadUrl(filename string, bucketname string) (string, error)
	GetIntDownloadUrl(filename string, bucketname string) (string, error)
}
