package dto

type StubFileTask struct {
	Id      string
	Objects struct {
		Edit string `json:"edit"`
	} `json:"objects"`
	Payload struct {
		Trace string `json:"trace"`
	} `json:"payload"`
}

type ResultTask struct {
	Id      string
	Type    string
	Status  string
	Objects map[string]string `json:"objects"`
	Payload Result            `json:"payload"`
}

type Result struct {
	Trace string
}

type WorkerRegister struct {
	Name    string         `json:"name"`
	Webhook string         `json:"webhook"`
	Scheme  map[string]any `json:"scheme"`
}
