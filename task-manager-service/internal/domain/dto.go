package domain

import (
	"github.com/google/uuid"
)

type Task struct {
	Id      uuid.UUID
	Type    string
	Objects map[string]string
	Payload map[string]any
}
