package dto

type S3Object struct {
	Key string `json:"key"`
}

type S3Record struct {
	S3 struct {
		Object S3Object `json:"object"`
	} `json:"s3"`
}

type S3Event struct {
	Records []S3Record `json:"Records"`
}

type UploadedFilesMessage struct {
	File   string `json:"file"`
	TaskID string `json:"task_id"`
	Type   string `json:"type"`
}
