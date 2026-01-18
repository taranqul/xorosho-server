package dto

type VideoConverterTask struct {
	Id      string
	Objects struct {
		Edit string `json:"edit"`
	} `json:"objects"`
	Payload struct {
		Target string `json:"target"`
	} `json:"payload"`
}

type ResultTask struct {
	Id      string
	Type    string
	Status  string
	Objects map[string]string `json:"objects"`
}

type WorkerRegister struct {
	Name    string         `json:"name"`
	Webhook string         `json:"webhook"`
	Scheme  map[string]any `json:"scheme"`
}
