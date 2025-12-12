package dto

type Task struct {
	Id      string
	Type    string
	Status  string
	Objects map[string]string
	Payload map[string]any
}
