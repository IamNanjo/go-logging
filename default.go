package logging

import (
	"github.com/IamNanjo/go-logging/pkg/ansi"
)

// Default logger

// Default logger.
// Out has no prefix, others have colored loglevel.LogLevel as prefix.
var Default = NewLogger(LoggerConfig{
	OutPrefix:     ansi.ColoredText{Color: ansi.None, Text: ""},
	DebugPrefix:   ansi.ColoredText{Color: ansi.Purple, Text: "DEBUG"},
	OkPrefix:      ansi.ColoredText{Color: ansi.Green, Text: "OK"},
	PendingPrefix: ansi.ColoredText{Color: ansi.Cyan, Text: "PENDING"},
	InfoPrefix:    ansi.ColoredText{Color: ansi.Blue, Text: "INFO"},
	WarnPrefix:    ansi.ColoredText{Color: ansi.Yellow, Text: "WARNING"},
	ErrPrefix:     ansi.ColoredText{Color: ansi.Red, Text: "ERROR"},
})
