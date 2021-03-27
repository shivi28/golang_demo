package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(WeekStart(2018, 1))
	fmt.Println(WeekStart(2018, 2))
	fmt.Println(WeekStart(2019, 1))
	fmt.Println(WeekStart(2021, 12))
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}
