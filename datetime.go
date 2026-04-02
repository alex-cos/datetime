package datetime

import (
	"fmt"
	"time"
)

const (
	Day            = 24 * time.Hour
	lastMonthIndex = 11
)

func SimpleFormat(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

func FormatISO8601Millis(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05.000Z")
}

func ParseISO8601Millis(str string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", str)
	if err != nil {
		err = fmt.Errorf("failed to parse datetime '%s': %w", str, err)
	}
	return t, err
}

// MidNightUTC returns the beginning of day for the given date.
func MidNightUTC(date time.Time) time.Time {
	return date.UTC().Truncate(Day)
}

// IsMidNightUTC returns true if the given date and time is in the midnight hour.
func IsMidNightUTC(date time.Time) bool {
	return date.UTC().Truncate(time.Hour).Hour() == 0
}

// LastMonth returns the first day of the last month for the given date.
func LastMonth(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	month := (d.Month() - 1)
	month--
	if month < 0 {
		month += 12
		year--
	}
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
}

// ThisMonth returns the first day of the current month for the given date.
func ThisMonth(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	month := d.Month()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// NextMonth returns the first day of the next month for the given date.
func NextMonth(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	month := (d.Month() - 1)
	month++
	if month > lastMonthIndex {
		month -= 12
		year++
	}
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
}

// LastQuarter returns the first day of last quarter for the given date.
func LastQuarter(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	month := (d.Month() - 1)
	month -= month % 3
	month -= 3
	if month < 0 {
		month += 12
		year--
	}
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
}

// ThisQuarter returns the first day of the current quarter for the given date.
func ThisQuarter(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	month := (d.Month() - 1)
	month -= month % 3
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
}

// NextQuarter returns the first day of the next quarter for the given date.
func NextQuarter(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	month := (d.Month() - 1)
	month -= month % 3
	month += 3
	if month > lastMonthIndex {
		month -= 12
		year++
	}
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
}

// LastYear returns the first day of last year for the given date.
func LastYear(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year() - 1
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// ThisYear returns the first day of the current year for the given date.
func ThisYear(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year()
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// NextYear returns the first day of the next year for the given date.
func NextYear(t time.Time) time.Time {
	d := t.UTC()
	year := d.Year() + 1
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// IsBeginingMonth returns true if t correspond to the first day of a month.
func IsBeginingMonth(t time.Time) bool {
	return (t.Day() == 1)
}

// IsBeginingQuarter returns true if t correspond to the first day of a quarter.
func IsBeginingQuarter(t time.Time) bool {
	return (t.Day() == 1 &&
		(t.Month() == time.January ||
			t.Month() == time.April ||
			t.Month() == time.July ||
			t.Month() == time.October))
}

// IsBeginingYear returns true if t correspond to the first day of a year.
func IsBeginingYear(t time.Time) bool {
	return (t.Day() == 1 && t.Month() == time.January)
}
