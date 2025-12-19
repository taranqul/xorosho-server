package dto

type WorkerRegister struct {
	Name    string         `bson:"_id" json:"name"`
	Webhook string         `bson:"webhook" json:"webhook"`
	Scheme  map[string]any `bson:"scheme" json:"scheme"`
}

type TaskRequest struct {
	Type    string            `bson:"type"  json:"task_type"`
	Objects map[string]string `bson:"objects" json:"objects"`
	Payload map[string]any    `bson:"payload" json:"payload"`
}
