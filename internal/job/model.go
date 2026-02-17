package job

import "time"

type Status string

const (
	StatusPending Status = "PENDING"
	StatusRunning Status = "RUNNING"
	StatusCompleted Status = "COMPLETED"
	StatusFailed Status = "FAILED"
)

type Job struct {
	ID string
	URL string
	Status Status
	Error *string
	CreatedAt time.Time
	UpdatedAt time.Time
}