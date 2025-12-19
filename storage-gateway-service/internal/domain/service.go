package domain

type GatewayService struct {
	external_storage S3Compatibile
	internal_storage S3Compatibile
}

func NewGatewayService(external_storage S3Compatibile, internal_storage S3Compatibile) *GatewayService {
	return &GatewayService{
		external_storage,
		internal_storage,
	}
}

func (g *GatewayService) GetExtUploadUrl(filename string, bucketname string) (string, error) {
	url, err := g.external_storage.GetUploadUrl(filename, bucketname)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (g *GatewayService) GetExtDownloadUrl(filename string, bucketname string) (string, error) {
	url, err := g.external_storage.GetDownloadUrl(filename, bucketname)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (g *GatewayService) GetIntUploadUrl(filename string, bucketname string) (string, error) {
	url, err := g.internal_storage.GetUploadUrl(filename, bucketname)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (g *GatewayService) GetIntDownloadUrl(filename string, bucketname string) (string, error) {
	url, err := g.internal_storage.GetDownloadUrl(filename, bucketname)
	if err != nil {
		return "", err
	}
	return url, nil
}
