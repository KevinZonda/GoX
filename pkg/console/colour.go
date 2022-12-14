package console

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	cfg, needReset := p.ConsoleString()
	fmt.Printf(cfg+format, params...)
	if needReset {
		ResetColour()
	}
}

func (p PrintConfig) WriteLine(format string, params ...interface{}) {
	p.Write(format+"\n", params...)
}

func (p PrintConfig) ErrWrite(format string, params ...interface{}) {
	cfg, needReset := p.ConsoleString()
	fmt.Fprintf(os.Stderr, cfg+format, params...)
	if needReset {
		ResetColour()
	}
}

func (p PrintConfig) ErrWriteLine(format string, params ...interface{}) {
	p.ErrWrite(format+"\n", params...)
}

func ResetColour() {
	fmt.Printf(ResetColourSymbol)
}

func (p PrintConfig) ConsoleString() (cfg string, needReset bool) {
	var sb strings.Builder
	sb.WriteString(p.Foreground.Foreground())
	sb.WriteString(p.Background.Background())
	if p.Bold {
		sb.WriteString(Bold)
	}
	if p.Underline {
		sb.WriteString(Underline)
	}
	cfg = sb.String()
	needReset = p.Underline || p.Bold || p.Background != None || p.Foreground != None
	return
}
