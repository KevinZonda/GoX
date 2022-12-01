package timex

import "testing"

func TestTimer(t *testing.T) {
	timer := NewTimer()
	timer.Start()
	_ = 1 + 1
	timer.Stop()
	if timer.Duration() < 0 {
		t.Fatal("Expected timer.Duration() >= 0, got", timer.Duration())
	}
}
