package domain

type GatewayService struct {
	storage S3Compatibile
}

func NewGatewayService(storage S3Compatibile) *GatewayService {
	return &GatewayService{
		storage,
	}
}

func (g *GatewayService) GetUploadUrl(filename string, bucketname string) (string, error) {
	url, err := g.storage.GetUploadUrl(filename, bucketname)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (g *GatewayService) GetDownloadUrl(filename string, bucketname string) (string, error) {
	url, err := g.storage.GetDownloadUrl(filename, bucketname)
	if err != nil {
		return "", err
	}
	return url, nil
}
