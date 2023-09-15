package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CronExpression represents a cron expression.
type CronExpression struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
}

func ParseCronExpression(expr string) (*CronExpression, error) {
	parts := strings.Fields(expr)
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid cron expression: %s", expr)
	}

	return &CronExpression{
		Minute:     parts[0],
		Hour:       parts[1],
		DayOfMonth: parts[2],
		Month:      parts[3],
	}, nil
}

func (cronExpr *CronExpression) CalculateNextRun(currentTime time.Time) (time.Time, error) {
	currentMinute := currentTime.Minute()
	currentHour := currentTime.Hour()
	currentDayOfMonth := currentTime.Day()
	currentMonth := int(currentTime.Month())
	currentYear := currentTime.Year()

	inputMinute := parseCronField(cronExpr.Minute, currentMinute)
	inputHour := parseCronField(cronExpr.Hour, currentHour)
	inputDayOfMonth := parseCronField(cronExpr.DayOfMonth, currentDayOfMonth)
	inputMonth := parseCronField(cronExpr.Month, currentMonth)

	inputMonth = currentMonth
	if cronExpr.Month != "*" {
		parsedMonth, err := strconv.Atoi(cronExpr.Month)
		if err == nil {
			inputMonth = parsedMonth
		}
		if inputMonth < currentMonth {
			currentYear++
		} else if inputMonth == currentMonth && inputDayOfMonth < currentDayOfMonth {
			currentYear++
			inputHour, inputMinute = setHourMinute(cronExpr, inputHour, inputMinute)
		}
		if currentMonth != inputMonth {
			if cronExpr.DayOfMonth == "*" {
				inputDayOfMonth = 1
			}
			inputHour, inputMinute = setHourMinute(cronExpr, inputHour, inputMinute)
		}
	}

	if cronExpr.DayOfMonth != "*" {
		if currentMonth == inputMonth && inputDayOfMonth < currentDayOfMonth && cronExpr.Month == "*" {
			// Check for the next available day in the next month
			inputMonth++
			inputHour, inputMinute = setHourMinute(cronExpr, inputHour, inputMinute)
		} else if currentMonth == inputMonth {
			if inputDayOfMonth > currentDayOfMonth {
				inputHour, inputMinute = setHourMinute(cronExpr, inputHour, inputMinute)
			} else if inputDayOfMonth == currentDayOfMonth {
				if cronExpr.Hour == "*" && cronExpr.Minute == "*" {
					inputHour = currentHour
					inputMinute = inputMinute + 1

					if inputMinute == 60 {
						inputHour++
						inputMinute = 0
					}
				} else if cronExpr.Hour != "*" && cronExpr.Minute == "*" {
					if inputHour > currentHour {
						inputMinute = 0
					} else if (inputHour < currentHour && cronExpr.Month != "*") || (inputHour == currentHour && inputMinute < currentMinute) {
						currentYear++
					} else if inputHour < currentHour && cronExpr.Month == "*" {
						inputMonth++
						inputHour, inputMinute = setHourMinute(cronExpr, inputHour, inputMinute)
					}
				} else if cronExpr.Hour != "*" && cronExpr.Minute != "*" {
					if inputHour < currentHour || (inputHour == currentHour && inputMinute < currentMinute) {
						if cronExpr.Month == "*" {
							inputMonth++
						} else {
							currentYear++
						}
					}
				}
			}
		}
	}

	// Handle additional cases
	if cronExpr.DayOfMonth == "*" && cronExpr.Hour == "*" && cronExpr.Minute != "*" {
		if inputMinute <= currentMinute {
			inputHour++
			if inputHour > 23 {
				inputHour = 0
				inputDayOfMonth++
				if inputDayOfMonth > daysInMonth(currentYear, inputMonth) {
					inputDayOfMonth = 1
					inputMonth++
					if inputMonth > 12 {
						inputMonth = 1
						currentYear++
					}
				}
			}
		}
	} else if cronExpr.DayOfMonth == "*" && cronExpr.Hour != "*" && cronExpr.Minute == "*" {
		if inputHour < currentHour {
			inputDayOfMonth++
			if inputDayOfMonth > daysInMonth(currentYear, inputMonth) {
				inputDayOfMonth = 1
				inputMonth++
				if inputMonth > 12 {
					inputMonth = 1
					currentYear++
				}
			}
			inputMinute = 0
		} else if inputHour == currentHour {
			inputMinute = currentMinute + 1
		} else {
			inputMinute = 0
		}
	} else if cronExpr.DayOfMonth == "*" && cronExpr.Hour != "*" && cronExpr.Minute != "*" {
		if inputHour < currentHour || (inputHour == currentHour && inputMinute < currentMinute) {
			inputDayOfMonth++
		}
	}

	nextRun := time.Date(currentYear, time.Month(inputMonth), inputDayOfMonth, inputHour, inputMinute, 0, 0, currentTime.Location())
	return nextRun, nil
}

func parseCronField(field string, current int) int {
	if field == "*" {
		return current
	}
	parsedValue, err := strconv.Atoi(field)
	if err != nil {
		return current
	}
	return parsedValue
}

func setHourMinute(cronExpr *CronExpression, currentHour, currentMinute int) (int, int) {
	if cronExpr.Hour == "*" && cronExpr.Minute == "*" {
		return 0, 0
	} else if cronExpr.Hour != "*" && cronExpr.Minute == "*" {
		return currentHour, 0
	} else if cronExpr.Hour == "*" && cronExpr.Minute != "*" {
		return 0, currentMinute
	}
	return currentHour, currentMinute
}

// daysInMonth returns the number of days in a given month.
func daysInMonth(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
