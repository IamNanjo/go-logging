package formatting

import (
	"strings"

	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

// Message begins after IndentSize characters. Automatically increased to the longest prefix out of all loggers.
var IndentSize = 7

func IndentWithPrefixAndSuffix(t *timeprefix.TimePrefix, p ansi.ColoredText, input string, s ansi.ColoredText) string {
	var (
		result strings.Builder
		indent = "\n" + strings.Repeat(" ", IndentSize)
		last   = len(input) - 1
	)

	var prefix string
	var visiblePrefixLength int
	var estimatedLength = last

	if estimatedLength > 0 {
		result.Grow(estimatedLength)
	}

	timePrefix := t.Get()
	result.WriteString(timePrefix)
	visiblePrefixLength = len(timePrefix) + len(p.Text)
	prefix += p.String()
	estimatedLength += len(prefix)

	result.WriteString(LeftPad(prefix, visiblePrefixLength, IndentSize))

	for i, r := range input {
		if r == '\n' {
			if i < last {
				result.WriteString(indent)
				continue
			} else if i == last {
				result.WriteString(s.String())
			}
		}
		result.WriteRune(r)
	}

	return result.String()
}
