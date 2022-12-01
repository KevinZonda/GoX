package console

import (
	"strings"
)

type ColourText struct {
	PrintConfig
	Text string
}

func NewColourText(text string) ColourText {
	return ColourText{
		Text: text,
	}
}

func ColourfulText(text string, foreground Colour, background Colour) ColourText {
	return ColourText{
		PrintConfig: PrintConfig{
			Foreground: foreground,
			Background: background,
			Bold:       false,
			Underline:  false,
		},
		Text: text,
	}
}

func (ct ColourText) String() string {
	var sb strings.Builder
	cfg, needReset := ct.ConsoleString()
	sb.WriteString(cfg)
	sb.WriteString(ct.Text)
	if needReset {
		sb.WriteString(ResetColourSymbol)
	}
	return sb.String()
}
