package models

import (
	"time"
)

// CronRun represents a record of a cron job run.
type CronRun struct {
	ID         string
	RunAt      time.Time
	Expression string
}

func NewCronRun(id string, runAt time.Time, expression string) *CronRun {
	return &CronRun{
		ID:         id,
		RunAt:      runAt,
		Expression: expression,
	}
}
