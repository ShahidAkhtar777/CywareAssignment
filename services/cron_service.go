package services

import (
	"CywareAssignment/models"
	"CywareAssignment/repositories"
	"CywareAssignment/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
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

		fmt.Printf("Next Run At: %s\n", nextRun.Format("02-01-2006 15:04:05"))
		cronRun := models.NewCronRun(uuid.New().String(), nextRun, cronExpression)
		s.repository.AddRun(cronRun)

		return nextRun, nil
	}
}

func (s *CronService) GetRecentCronRuns() []*models.CronRun {
	return s.repository.GetRuns()
}
