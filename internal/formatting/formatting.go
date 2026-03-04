package formatting

import "strings"

func IndentNewLines(input string, prefixWidth int) string {
	var (
		newValue strings.Builder
		indent   = "\n" + strings.Repeat(" ", prefixWidth)
		last     = len(input) - 1
	)
	newValue.Grow(last)
	for i, r := range input {
		newValue.WriteRune(r)
		if r == '\n' && i < last {
			newValue.WriteString(indent)
		}
	}
	return newValue.String()
}
