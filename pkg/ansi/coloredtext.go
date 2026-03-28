package ansi

import (
	"os"

	"golang.org/x/term"
)

type ColoredText struct {
	Color Color
	Text  string
}

var useColor = term.IsTerminal(int(os.Stdout.Fd())) && os.Getenv("NO_COLOR") == "" && os.Getenv("TERM") != "dumb"

// Output string with ANSI escape codes unless output is not a terminal,
// user has NO_COLOR environment variable set or
// TERM environment variable is set to dumb.
func (ct *ColoredText) String() string {
	if !useColor || ct.Color == None {
		return ct.Text
	}
	return string(ct.Color) + ct.Text + string(Reset)
}
