package formatting

import "strings"

func IndentNewLines(input string, prefixWidth int) string {
	return strings.ReplaceAll(
		input,
		"\n",
		"\n"+strings.Repeat(" ", prefixWidth),
	)
}
