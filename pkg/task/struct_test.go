package task

import "testing"

func TestTask(t *testing.T) {
	m := New[Void](func() Void {
		println("Hello, World!")
		return VoidResult
	})
	Async(m)
	m.Wait()
}
