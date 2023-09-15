package repositories

import (
	"CywareAssignment/models"
	"sync"
)

// CronRepository represents a repository for storing and retrieving CronRun records.
type CronRepository struct {
	mu   sync.RWMutex
	runs []*models.CronRun
}

func NewCronRepository() *CronRepository {
	return &CronRepository{}
}

func (r *CronRepository) AddRun(run *models.CronRun) {
	r.runs = append(r.runs, run)
}

func (r *CronRepository) GetRuns() []*models.CronRun {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.runs
}
