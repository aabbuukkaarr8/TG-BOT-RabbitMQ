package tg_bot

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseTime(timeStr string) time.Time {
	now := time.Now().In(time.Local)

	if strings.Contains(timeStr, "минут") {
		// "через 30 минут"
		minutes := extractNumber(timeStr)
		newTime := now.Add(time.Duration(minutes) * time.Minute)
		return newTime
	}

	if strings.Contains(timeStr, "час") {
		// "через 2 часа"
		hours := extractNumber(timeStr)
		newTime := now.Add(time.Duration(hours) * time.Hour)
		return newTime
	}

	return time.Time{} // не распознано
}
func extractNumber(s string) int {
	// Ищем цифры в строке
	re := regexp.MustCompile(`\d+`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return 0
	}

	num, _ := strconv.Atoi(matches[0])
	return num
}
