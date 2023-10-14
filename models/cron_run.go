package models

import (
	"time"
)

// CronRun represents a record of a cron job run.
type CronRun struct {
	ID         uint `gorm:"primary_key"`
	RunAt      time.Time
	Expression string
}

func NewCronRun(runAt time.Time, expression string) *CronRun {
	return &CronRun{
		RunAt:      runAt,
		Expression: expression,
	}
}
