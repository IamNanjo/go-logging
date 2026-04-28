package formatting

import (
	"strings"

	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

var (
	// Message begins after IndentSize characters. Automatically increased when creating new loggers.
	IndentSize int
	// Length of time longest time prefix. Automatically increased when creating new loggers.
	TimeLen         int
	columnSeparator = " │ "
)

// Column separator length in visible characters
const ColumnSeparatorLen = 3

func IndentWithPrefixAndSuffix(
	t *timeprefix.TimePrefix,
	prefix ansi.ColoredText,
	input string,
	suffix ansi.ColoredText,
) string {
	var (
		result strings.Builder
		last   = len(input) - 1
	)
	if last > 0 {
		result.Grow(last)
	}

	// Time prefix

	timePrefix := t.Get()
	timePadding := GetPadding(len(timePrefix), TimeLen)

	result.WriteString(timePrefix)
	result.Write(timePadding)

	if TimeLen != 0 {
		result.WriteString(columnSeparator)
	}

	// Prefix

	padding := GetPadding(TimeLen+len(prefix.Text), IndentSize)

	result.Write(padding)
	result.WriteString(prefix.String())
	result.WriteString(columnSeparator)

	// Multiline handling

	var indent = append(GetPadding(0, IndentSize), columnSeparator...)
	if TimeLen != 0 {
		copy(indent[TimeLen-ColumnSeparatorLen:TimeLen], []byte(columnSeparator))
	}

	for i, r := range input {
		if r != '\n' {
			result.WriteRune(r)
			continue
		}

		if i < last {
			result.WriteRune(r)
			result.Write(indent)
		} else if i == last {
			result.WriteString(suffix.String())
			result.WriteRune(r)
		}
	}

	return result.String()
}
