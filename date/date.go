package date

import (
	"time"
)

func GetMondayOfWeek(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	weekday := t.Weekday()
	offset := weekday - time.Monday
	if offset == -1 {
		offset += 7
	}
	return t.AddDate(0, 0, -1*int(offset))
}

func GetMondyOfThisWeek() time.Time {
	return GetMondayOfWeek(time.Now())
}

func GetFridayOfWeek(t time.Time) time.Time {
	return GetMondayOfWeek(t).AddDate(0, 0, 4)
}

func GetFridayOfThisWeek() time.Time {
	return GetMondyOfThisWeek().AddDate(0, 0, 4)
}
