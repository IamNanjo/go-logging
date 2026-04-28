package logging

import (
	"time"

	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
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

var defaultT *Logger

// Returns default logger with RFC3339 timestamps.
func DefaultT() *Logger {
	if defaultT != nil {
		return defaultT
	}

	c := Default.Config()
	c.Time = &timeprefix.TimePrefix{
		Format: time.RFC3339,
	}

	defaultT = NewLogger(*c)
	return defaultT
}
