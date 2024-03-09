package calendar_test

import (
	"github.com/Jumpaku/gkoku/calendar"
	"time"
)

func DateFromTime(t time.Time) calendar.Date {
	return calendar.YyyyMmDd(t.Year(), calendar.Month(t.Month()), t.Day())
}
func ToTime(y int, m calendar.Month, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}