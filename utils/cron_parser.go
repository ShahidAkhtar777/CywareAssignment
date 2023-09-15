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

	inputMinute := currentMinute
	if cronExpr.Minute != "*" {
		parsedMinute, err := strconv.Atoi(cronExpr.Minute)
		if err == nil {
			inputMinute = parsedMinute
		}
	}

	inputHour := currentHour
	if cronExpr.Hour != "*" {
		parsedHour, err := strconv.Atoi(cronExpr.Hour)
		if err == nil {
			inputHour = parsedHour
		}
	}

	inputDayOfMonth := currentDayOfMonth
	if cronExpr.DayOfMonth != "*" {
		parsedDay, err := strconv.Atoi(cronExpr.DayOfMonth)
		if err == nil {
			inputDayOfMonth = parsedDay
		}
	}

	inputMonth := currentMonth
	if cronExpr.Month != "*" {
		parsedMonth, err := strconv.Atoi(cronExpr.Month)
		if err == nil {
			inputMonth = parsedMonth
		}
	}

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
			if cronExpr.Hour == "*" && cronExpr.Minute == "*" {
				inputHour = 0
				inputMinute = 0
			} else if cronExpr.Hour != "*" && cronExpr.Minute == "*" {
				inputMinute = 0
			}
		}
		if currentMonth != inputMonth {
			if cronExpr.DayOfMonth == "*" {
				inputDayOfMonth = 1
			}
			if cronExpr.Hour == "*" && cronExpr.Minute == "*" {
				inputHour = 0
				inputMinute = 0
			} else if cronExpr.Hour != "*" && cronExpr.Minute == "*" {
				inputMinute = 0
			}
		}
	}

	if cronExpr.DayOfMonth != "*" {
		if currentMonth == inputMonth && inputDayOfMonth < currentDayOfMonth && cronExpr.Month == "*" {
			// Check for the next available day in the next month
			inputMonth++
			if cronExpr.Hour == "*" && cronExpr.Minute == "*" {
				inputHour = 0
				inputMinute = 0
			} else if cronExpr.Hour != "*" && cronExpr.Minute == "*" {
				inputMinute = 0
			}
		} else if currentMonth == inputMonth {
			if inputDayOfMonth > currentDayOfMonth {
				if cronExpr.Hour == "*" && cronExpr.Minute == "*" {
					inputHour = 0
					inputMinute = 0
				} else if cronExpr.Hour != "*" && cronExpr.Minute == "*" {
					inputMinute = 0
				}
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

// daysInMonth returns the number of days in a given month.
func daysInMonth(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
