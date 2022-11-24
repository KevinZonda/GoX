package console

import "testing"

func TestColour(t *testing.T) {
	Red.AsForeground().WithBackground(Green).Write("Hello, world!")
}

func TestWhiteColour(t *testing.T) {
	Cyan.AsForeground().WithBackground(White).Write("Hello, world!")
}
