package timex

import "time"

type Timer struct {
	cur *time.Time
	dur time.Duration
}

func NewTimer() *Timer {
	return &Timer{
		cur: nil,
		dur: -1,
	}
}

func (t *Timer) Start() {
	now := time.Now()
	t.cur = &now
}

func (t *Timer) Stop() {
	if t.cur == nil {
		return
	}
	t.dur = time.Since(*t.cur)
	t.cur = nil
	return
}

func (t *Timer) Reset() {
	t.cur = nil
}

func (t *Timer) Duration() time.Duration {
	return t.dur
}

func (t *Timer) IsStarted() bool {
	return t.cur != nil
}

func (t *Timer) IsStopped() bool {
	return t.cur == nil
}

func (t *Timer) Elapsed() time.Duration {
	if t.cur == nil {
		return -1
	}
	return time.Since(*t.cur)
}
