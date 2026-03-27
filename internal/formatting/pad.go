package formatting

import (
	"strings"
)

// Add spaces until input is at least as long as size
func LeftPad(input string, visibleLength int, size int) string {
	padding := size - visibleLength
	if padding <= 0 {
		return input
	}

	return strings.Repeat(" ", padding) + input
}
