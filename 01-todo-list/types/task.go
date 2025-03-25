package types

import "time"

type Task struct {
	ID        uint
	Task      string
	CreatedAt time.Time
	IsComplete bool
} 