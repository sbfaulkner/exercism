package meetup

import (
	"time"
)

const testVersion = 3

// WeekSchedule represents how to schedule a meetup
type WeekSchedule int

// WeekSchedule values
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the date of the scheduled meetup
func Day(s WeekSchedule, w time.Weekday, m time.Month, y int) (day int) {
	switch s {
	case Teenth:
		date := time.Date(y, m, 13, 0, 0, 0, 0, time.UTC)
		day = date.AddDate(0, 0, daysUntilWeekday(date, w)).Day()
	case Last:
		date := time.Date(y, m+1, 1, 0, 0, 0, 0, time.UTC)
		date = date.AddDate(0, 0, -1)
		day = date.AddDate(0, 0, -daysAfterWeekday(date, w)).Day()
	default:
		date := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
		day = date.AddDate(0, 0, daysUntilWeekday(date, w)+int(s)*7).Day()
	}

	return
}

func daysUntilWeekday(date time.Time, w time.Weekday) int {
	return (7 + int(w-date.Weekday())) % 7
}

func daysAfterWeekday(date time.Time, w time.Weekday) int {
	return (7 + int(date.Weekday()-w)) % 7
}
