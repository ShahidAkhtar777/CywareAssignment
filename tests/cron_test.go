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

	// Test monthly cron at 15th of every month at 4:30am [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("30 4 15 *")
	expectedNextRun = time.Date(2023, time.October, 15, 4, 30, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test monthly cron at 15th of every month at 10:30am [Time has not passed for today]
	cronExpr, _ = utils.ParseCronExpression("30 10 15 *")
	expectedNextRun = time.Date(2023, time.September, 15, 10, 30, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test past day for current month -> give next yr [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("* * 4 9")
	expectedNextRun = time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test Upcoming month day time  [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("15 15 20 12")
	expectedNextRun = time.Date(2023, time.December, 20, 15, 15, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test past time in current day ->Should give next yr [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("12 6 15 9")
	expectedNextRun = time.Date(2024, time.September, 15, 6, 12, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test every day in current month at 6:12 am [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("12 6 * 9")
	expectedNextRun = time.Date(2023, time.September, 16, 6, 12, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}

	// Test every 2nd day of every month at 5am  [Time has passed for today]
	cronExpr, _ = utils.ParseCronExpression("* 5 2 *")
	expectedNextRun = time.Date(2023, time.October, 2, 5, 0, 0, 0, time.UTC)
	nextRun, _ = cronExpr.CalculateNextRun(currentTime)
	if !nextRun.Equal(expectedNextRun) {
		t.Errorf("Expected next run: %s, got: %s", expectedNextRun, nextRun)
	}
}
