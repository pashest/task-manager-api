package model

import (
	"github.com/google/uuid"
	"time"
)

type TaskStatus string

const (
	StatusCreated    TaskStatus = "CREATED"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusCompleted  TaskStatus = "COMPLETED"
	StatusCanceled   TaskStatus = "CANCELED"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
}
