package formatting

import (
	"strings"

	"github.com/IamNanjo/go-logging/internal/common"
	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

// Message begins after IndentSize characters. Automatically increased to the longest prefix out of all loggers.
var (
	IndentSize int
	TimeLen    int
)

func IndentWithPrefixAndSuffix(
	t *timeprefix.TimePrefix,
	prefix ansi.ColoredText,
	input string,
	suffix ansi.ColoredText,
) string {
	var (
		result strings.Builder
		indent = append(GetPadding(0, IndentSize), common.ColumnSeparator...)
	)

	var last = len(input) - 1
	if last > 0 {
		result.Grow(last)
	}

	timePrefix := t.Get()
	result.WriteString(timePrefix)

	padding := GetPadding(len(timePrefix)+len(prefix.Text), IndentSize)
	if TimeLen > common.ColumnSeparatorLen && timePrefix == "" {
		copy(padding[TimeLen-common.ColumnSeparatorLen:TimeLen], common.ColumnSeparator)
	}
	result.Write(padding)
	result.WriteString(prefix.String())
	result.Write(common.ColumnSeparator)

	if TimeLen > common.ColumnSeparatorLen {
		copy(indent[TimeLen-common.ColumnSeparatorLen:TimeLen], common.ColumnSeparator)
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
