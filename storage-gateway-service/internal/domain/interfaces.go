package domain

type S3Compatibile interface {
	GetUploadUrl(filename string, bucketname string) (string, error)
	GetDownloadUrl(filename string, bucketname string) (string, error)
}
