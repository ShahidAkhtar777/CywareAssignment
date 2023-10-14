package services

import (
	"CywareAssignment/models"
	"CywareAssignment/repositories"
	"CywareAssignment/utils"
	"errors"
	"fmt"
	"time"
)

type CronService struct {
	repository *repositories.CronRepository
}

func NewCronService(repository *repositories.CronRepository) *CronService {
	return &CronService{
		repository: repository,
	}
}

func (s *CronService) CalculateNextRun(cronExpression string) (time.Time, error) {
	cronExpr, err := utils.ParseCronExpression(cronExpression)
	if err != nil {
		return time.Time{}, errors.New("error parsing cron expression")
	} else {
		now := time.Now()
		nextRun, err := cronExpr.CalculateNextRun(now)
		if err != nil {
			return time.Time{}, errors.New("error calculating next run")
		}

		cronRun := models.NewCronRun(nextRun, cronExpression)
		if err := utils.DB.Create(&cronRun).Error; err != nil {
			return time.Time{}, err
		}

		fmt.Printf("Next Run At: %s\n", nextRun.Format("02-01-2006 15:04:05"))

		return nextRun, nil
	}
}

func (s *CronService) GetRecentCronRuns() []*models.CronRun {
	var runs []*models.CronRun
	if err := utils.DB.Find(&runs).Error; err != nil {
		return nil
	}
	return runs
}
