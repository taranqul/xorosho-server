package domain

type Task struct {
	Id      string
	Type    string
	Objects map[string]string
	Payload map[string]any
}

type TaskInRepository struct {
	Id      string            `bson:"_id" json:"id"`
	Status  string            `bson:"status" json:"status"`
	Type    string            `bson:"type"  json:"type"`
	Objects map[string]string `bson:"objects" json:"objects"`
	Payload map[string]any    `bson:"payload" json:"payload"`
}

type UploadedFilesMessage struct {
	File   string `json:"file"`
	TaskID string `json:"task_id"`
	Type   string `json:"type"`
}
