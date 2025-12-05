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

type TaskInRepository struct {
	Id      string            `bson:"_id"`
	Status  string            `bson:"status"`
	Type    string            `bson:"type"`
	Objects map[string]string `bson:"objects"`
	Payload map[string]any    `bson:"payload"`
}
