package dto

type WorkerRegister struct {
	Name    string         `bson:"_id" json:"name"`
	Webhook string         `bson:"webhook" json:"webhook"`
	Scheme  map[string]any `bson:"scheme" json:"scheme"`
}
