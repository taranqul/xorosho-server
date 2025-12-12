package domain

type Task struct {
	Id      string
	Type    string
	Objects map[string]string
	Payload map[string]any
}

type TaskInRepository struct {
	Id      string            `bson:"_id"`
	Status  string            `bson:"status"`
	Type    string            `bson:"type"`
	Objects map[string]string `bson:"objects"`
	Payload map[string]any    `bson:"payload"`
}

type UploadedFilesMessage struct {
	File   string `json:"file"`
	TaskID string `json:"task_id"`
	Type   string `json:"type"`
}
