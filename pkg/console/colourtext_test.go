package console

import (
	"fmt"
	"testing"
)

func TestColourText(t *testing.T) {
	ct := ColourText{Text: "Hello, world!"}
	ct.Foreground = Red
	ct.Background = Green
	ct.Bold = true
	ct.Underline = true
	fmt.Print(ct)
}
