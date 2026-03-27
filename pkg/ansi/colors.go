// Contains ANSI escape codes for different (high intensity) colors.
package ansi

import "strings"

type Color string

const (
	None         = "" // No color
	Reset  Color = "\033[0;0m"
	Black  Color = "\033[0;90m"
	Red    Color = "\033[0;91m"
	Green  Color = "\033[0;92m"
	Yellow Color = "\033[0;93m"
	Blue   Color = "\033[0;94m"
	Purple Color = "\033[0;95m"
	Cyan   Color = "\033[0;36m"
	White  Color = "\033[0;97m"
)

// Replace colors from string to accurately compare string lengths.
var Uncolor = strings.NewReplacer(
	string(Reset), "",
	string(Black), "",
	string(Red), "",
	string(Green), "",
	string(Yellow), "",
	string(Blue), "",
	string(Purple), "",
	string(Cyan), "",
	string(White), "",
)
