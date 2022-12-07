package timex

import "time"

const Day time.Duration = 24 * time.Hour

func FromSecond(sec int) time.Duration {
	return From(sec, time.Second)
}

func FromHour(sec int) time.Duration {
	return From(sec, time.Hour)
}

func FromDay(sec int) time.Duration {
	return From(sec, Day)
}

func FromMillisecond(sec int) time.Duration {
	return From(sec, time.Millisecond)
}

func FromMicrosecond(sec int) time.Duration {
	return From(sec, time.Microsecond)
}

func From(v int, unit time.Duration) time.Duration {
	return time.Duration(v) * unit
}

func Between(left, mid, right time.Duration) time.Duration {
	if left > right {
		return left - mid
	}
	return mid - left
}

// IsBetween returns true if the time is between the two times.
func IsBetween(left, mid, right time.Time) bool {
	return mid.After(left) && mid.Before(right)

}

// IsBetweenOrEqual returns true if the time is between the two times or equal to one of them.
func IsBetweenOrEqual(left, mid, right time.Time) bool {
	return mid.After(left) && mid.Before(right) || mid.Equal(left) || mid.Equal(right)
}

// NowIn returns the current time in the given location.
func NowIn(location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}
