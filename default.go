package logging

import (
	"github.com/IamNanjo/go-logging/pkg/ansi"
)

// Default logger

var defaultLogger = NewLogger(LoggerConfig{
	OutPrefix:     ansi.ColoredText{Color: ansi.None, Text: "OUT "},
	DebugPrefix:   ansi.ColoredText{Color: ansi.Purple, Text: "DEBUG "},
	OkPrefix:      ansi.ColoredText{Color: ansi.Green, Text: "OK "},
	PendingPrefix: ansi.ColoredText{Color: ansi.Cyan, Text: "PENDING "},
	InfoPrefix:    ansi.ColoredText{Color: ansi.Blue, Text: "INFO "},
	WarnPrefix:    ansi.ColoredText{Color: ansi.Yellow, Text: "WARNING "},
	ErrPrefix:     ansi.ColoredText{Color: ansi.Red, Text: "ERROR "},
})

func Out(format string, args ...any) (int, error) {
	return defaultLogger.Out(format, args)
}
func Debug(format string, args ...any) (int, error) {
	return defaultLogger.Debug(format, args)
}
func Ok(format string, args ...any) (int, error) {
	return defaultLogger.Ok(format, args)
}
func Pending(format string, args ...any) (int, error) {
	return defaultLogger.Pending(format, args)
}
func Info(format string, args ...any) (int, error) {
	return defaultLogger.Info(format, args)
}
func Warn(format string, args ...any) (int, error) {
	return defaultLogger.Warn(format, args)
}
func Err(format string, args ...any) (int, error) {
	return defaultLogger.Err(format, args)
}
func Fatal(format string, args ...any) {
	defaultLogger.Fatal(format, args)
}
