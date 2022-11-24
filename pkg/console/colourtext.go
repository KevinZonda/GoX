package console

import (
	"strings"
)

type ColourText struct {
	PrintConfig
	Text string
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
