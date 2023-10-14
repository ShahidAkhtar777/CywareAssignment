package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidCronExpression(expr string) bool {
	// Currently I have implemented only integers and wildcards but pattern can take step and range values.
	pattern := `^(?:[0-5]?[0-9]|\*)\s(?:[01]?[0-9]|2[0-3]|\*)\s(?:[1-9]|[12][0-9]|3[01]|\*)\s(?:[1-9]|1[0-2]|\*)$`
	match, err := regexp.MatchString(pattern, expr)
	if err != nil {
		return false
	}

	// Validate individual fields
	fields := strings.Fields(expr)
	if !IsValidCronField(fields[0], 0, 59) {
		return false
	}
	if !IsValidCronField(fields[1], 0, 23) {
		return false
	}
	if !IsValidCronField(fields[2], 1, 31) {
		return false
	}
	if !IsValidCronField(fields[3], 1, 12) {
		return false
	}

	return match
}

func IsValidCronField(field string, min int, max int) bool {
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
