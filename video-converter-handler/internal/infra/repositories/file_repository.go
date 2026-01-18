package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type FileRepository struct {
	logger *zap.Logger
}

func NewFileRepository(logger *zap.Logger) *FileRepository {
	return &FileRepository{
		logger: logger,
	}
}

func (f *FileRepository) GetFile(name string) (io.ReadCloser, error) {
	apiURL := "http://storage-gateway-service:8080/storage/internal/downloadUrl"

	params := url.Values{}
	params.Add("filename", name)
	params.Add("bucketname", "upload")

	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		f.logger.Sugar().Errorf("Failed to get download URL: %s", string(body))
		return nil, err
	}

	var presignedURL string
	err = json.NewDecoder(resp.Body).Decode(&presignedURL)
	if err != nil {
		f.logger.Sugar().Errorf("Got presigned URL:", presignedURL)
		return nil, err
	}
	f.logger.Sugar().Infof("Got presigned URL:", presignedURL)
	fileResp, err := http.Get(presignedURL)
	if err != nil {
		return nil, err
	}

	if fileResp.StatusCode != http.StatusOK {
		defer fileResp.Body.Close()
		f.logger.Sugar().Errorf("Failed to get file")
		return nil, err
	}

	return fileResp.Body, err
}

func (f *FileRepository) UploadFile(name string, data io.Reader) error {
	apiURL := "http://storage-gateway-service:8080/storage/internal/uploadUrl"
	params := url.Values{}
	params.Add("filename", name)
	params.Add("bucketname", "results")
	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("failed to get presigned URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		f.logger.Sugar().Errorf("Failed to get presigned URL: %s", string(body))
		return fmt.Errorf("failed to get presigned URL: %s", resp.Status)
	}

	var presignedURL string
	err = json.NewDecoder(resp.Body).Decode(&presignedURL)
	if err != nil {
		return fmt.Errorf("failed to decode presigned URL: %w", err)
	}

	f.logger.Sugar().Infof("Got presigned URL: %s", presignedURL)

	req, err := http.NewRequest(http.MethodPut, presignedURL, data)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %w", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	uploadResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode != http.StatusOK && uploadResp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(uploadResp.Body)
		f.logger.Sugar().Errorf("Failed to upload file: %s", string(body))
		return fmt.Errorf("failed to upload file: %s", uploadResp.Status)
	}

	f.logger.Sugar().Infof("File uploaded successfully: %s", name)
	return nil
}
