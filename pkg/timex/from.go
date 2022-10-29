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
