package console

import "testing"

func TestColour(t *testing.T) {
	Red.AsForeground().WithBackground(Green).Write("Hello, world!")
}
