package dto

type StubNoFileTask struct {
	Id      string
	Payload struct {
		Trace string `json:"trace"`
	} `json:"payload"`
}
type ResultTask struct {
	Id      string
	Type    string
	Status  string
	Objects map[string]string
	Payload Result `json:"payload"`
}

type Result struct {
	Trace string
}
