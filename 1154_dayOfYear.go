package main

import (
	"time"
)

func dayOfYear(date string) int {

	now, _ := time.ParseInLocation("2006-01-02", date, time.Local)

	return now.YearDay()
}
