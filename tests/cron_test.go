package tests

import (
	"CywareAssignment/utils"
	"testing"
	"time"
)

func TestCalculateNextRun(t *testing.T) {
	currentTime := time.Date(2023, time.September, 15, 7, 33, 0, 0, time.UTC)

	// Test a daily cron job at 10:15 AM [Time not passed for today]
	cronExpr, _ := utils.ParseCronExpression("15 10 * *")
	expectedNextRun := time.Date(2023, time.September, 15, 10, 15, 0, 0, time.UTC)
	nextRun, _ := cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Hour specified rest wildcard [Time not passed for today]
	cronExpr, _ = utils.ParseCronExpression("* 8 * *")
	expectedNextRun = time.Date(2023, time.September, 15, 8, 0, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Hour specified rest wildcard [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("* 4 * *")
	expectedNextRun = time.Date(2023, time.September, 16, 4, 0, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test Upcoming Time
	cronExpr, _ = utils.ParseCronExpression("* * 10 10")
	expectedNextRun = time.Date(2023, time.October, 10, 0, 0, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test next 0th minute at any hour [It should give next hour]
	cronExpr, _ = utils.ParseCronExpression("0 * * *")
	expectedNextRun = time.Date(2023, time.September, 15, 8, 0, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test 40th minute at any hour [It should give same hour and 40th minute]
	cronExpr, _ = utils.ParseCronExpression("40 * * *")
	expectedNextRun = time.Date(2023, time.September, 15, 7, 40, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}
}
