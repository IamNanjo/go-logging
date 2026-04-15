// Initialize logging level by reading case-insensitive environment variable LOGLEVEL.
// Valid values ordered from least verbose to most verbose are CRITICAL, ERROR, WARN, INFO, DEBUG.
// Level can also be set programmatically by editing loglevel.Level to any of the supported values.
package loglevel

import (
	"os"
	"strings"
)

type LogLevel uint8

//go:generate go tool stringer -type=LogLevel
const (
	SILENT LogLevel = iota
	CRITICAL
	ERROR
	WARN
	INFO
	DEBUG
)

var Level LogLevel

// Parse LogLevel from case-insensitive level.
func FromString(level string) LogLevel {
	switch strings.ToUpper(level) {
	case "CRITICAL":
		return CRITICAL
	case "ERROR":
		return ERROR
	case "WARN":
		return WARN
	case "INFO":
		return INFO
	case "DEBUG":
		return DEBUG
	default:
		return INFO
	}
}

func init() {
	Level = FromString(os.Getenv("LOGLEVEL"))
}
