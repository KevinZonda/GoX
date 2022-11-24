package console

import (
	"fmt"
	"strconv"
)

type Colour int

const (
	Black  Colour = 0
	Red    Colour = 1
	Green  Colour = 2
	Yellow Colour = 3
	Blue   Colour = 4
	Purple Colour = 5
	Cyan   Colour = 6
	Gray   Colour = 7
	White  Colour = 67
	None   Colour = -1
)

const ResetColourSymbol = "\033[0m"

const (
	Bold      = "\033[1m"
	Underline = "\033[4m"
)

func (c Colour) Foreground() string {
	if c == None {
		return ""
	}
	return "\033[" + strconv.Itoa(int(30+c)) + "m"
}

func (c Colour) Background() string {
	if c == None {
		return ""
	}
	return "\033[" + strconv.Itoa(int(40+c)) + "m"
}

type PrintConfig struct {
	Foreground Colour
	Background Colour
	Bold       bool
	Underline  bool
}

func dftPrintConfig() PrintConfig {
	return PrintConfig{
		Foreground: None,
		Background: None,
		Bold:       false,
		Underline:  false,
	}
}

func (c Colour) AsForeground() PrintConfig {
	p := dftPrintConfig()
	p.Foreground = c
	return p
}

func (c Colour) AsBackground() PrintConfig {
	p := dftPrintConfig()
	p.Background = c
	return p
}

func (p PrintConfig) WithForeground(c Colour) PrintConfig {
	p.Foreground = c
	return p
}

func (p PrintConfig) WithBackground(c Colour) PrintConfig {
	p.Background = c
	return p
}

func (p PrintConfig) WithBold(b bool) PrintConfig {
	p.Bold = b
	return p
}

func (p PrintConfig) WithUnderline(u bool) PrintConfig {
	p.Underline = u
	return p
}

func (p PrintConfig) Write(format string, params ...interface{}) {
	fmt.Printf(p.Foreground.Foreground())
	fmt.Printf(p.Background.Background())
	if p.Bold {
		fmt.Printf(Bold)
	}
	if p.Underline {
		fmt.Printf(Underline)
	}
	fmt.Printf(format, params...)
	if p.Underline == false && p.Bold == false && p.Background == None && p.Foreground == None {
		return
	}
	fmt.Printf(ResetColourSymbol)
}

func ResetColour() {
	fmt.Printf(ResetColourSymbol)
}
