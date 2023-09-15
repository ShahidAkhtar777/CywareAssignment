package main

import (
	"CywareAssignment/repositories"
	"CywareAssignment/services"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Cron CLI")
	fmt.Println("Type 'exit' to quit the application")

	// Initialize the repository and service
	cronRepository := repositories.NewCronRepository()
	cronService := services.NewCronService(cronRepository)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> Enter Expression (e.g., '* * * *'), 'fetch', or 'exit': ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		} else if strings.ToLower(input) == "fetch" {
			displayRecentCronRuns(cronService)
		} else {
			input = strings.TrimSpace(input)

			if isValidCronExpression(input) {
				_, err := cronService.CalculateNextRun(input)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Invalid input. Please enter a valid cron expression or 'fetch'.")
			}
		}
	}

	fmt.Println("Goodbye! ~ Interview Candidate: Shahid Akhtar")
}

func displayRecentCronRuns(cronService *services.CronService) {
	recentCronRuns := cronService.GetRecentCronRuns()
	fmt.Println("\nRecent Cron Job Runs:")
	for _, run := range recentCronRuns {
		fmt.Printf("ID: %s, Next Run At: %s, Cron Expression: %s\n", run.ID, run.RunAt.Format("02-01-2006 15:04:05"), run.Expression)
	}
}
func isValidCronExpression(expr string) bool {
	// Currently i have implemented only integers and wildcards but pattern can take step and range values.
	pattern := `^(?:[0-5]?[0-9]|\*)\s(?:[01]?[0-9]|2[0-3]|\*)\s(?:[1-9]|[12][0-9]|3[01]|\*)\s(?:[1-9]|1[0-2]|\*)$`
	match, err := regexp.MatchString(pattern, expr)
	if err != nil {
		return false
	}

	// Validate individual fields
	fields := strings.Fields(expr)
	if !isValidCronField(fields[0], 0, 59) {
		return false
	}
	if !isValidCronField(fields[1], 0, 23) {
		return false
	}
	if !isValidCronField(fields[2], 1, 31) {
		return false
	}
	if !isValidCronField(fields[3], 1, 12) {
		return false
	}

	return match
}

func isValidCronField(field string, min int, max int) bool {
	values := strings.Split(field, ",")
	for _, value := range values {
		if value == "*" {
			continue
		}
		if strings.Contains(value, "-") {
			rangeParts := strings.Split(value, "-")
			if len(rangeParts) != 2 {
				return false
			}
			start, err := strconv.Atoi(rangeParts[0])
			end, err2 := strconv.Atoi(rangeParts[1])
			if err != nil || err2 != nil || start < min || end > max || start > end {
				return false
			}
		} else if strings.Contains(value, "/") {
			stepParts := strings.Split(value, "/")
			if len(stepParts) != 2 {
				return false
			}
			step, err := strconv.Atoi(stepParts[1])
			if err != nil || step <= 0 || step > max {
				return false
			}
		} else {
			val, err := strconv.Atoi(value)
			if err != nil || val < min || val > max {
				return false
			}
		}
	}
	return true
}
