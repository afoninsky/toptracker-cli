package main

import (
	"time"
)

func workingDaysInRange(start, end time.Time) int {
	var total = 0
	for d := start; d.Month() == end.Month() && d.Day() <= end.Day(); d = d.AddDate(0, 0, 1) {
		dayStr := d.Format("2006-01-02")
		_, isHoliday := holidays[dayStr]
		if isHoliday {
			continue
		}
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			continue
		}
		total = total + 1
	}
	return total
}

// how previous working days were done
func effiency(shoudWorkSeconds, realWorkSeconds int) int {
	return realWorkSeconds * 100 / shoudWorkSeconds
}
