package repositories

import (
	"CywareAssignment/models"
	"github.com/jinzhu/gorm"
)

type CronRepository struct {
	db *gorm.DB
}

func NewCronRepository(db *gorm.DB) *CronRepository {
	return &CronRepository{db: db}
}

func (r *CronRepository) AddRun(run *models.CronRun) error {
	return r.db.Create(run).Error
}

func (r *CronRepository) GetRuns() ([]models.CronRun, error) {
	var runs []models.CronRun
	if err := r.db.Find(&runs).Error; err != nil {
		return nil, err
	}
	return runs, nil
}
