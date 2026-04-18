package logging

import (
	"github.com/IamNanjo/go-logging/internal/formatting"
	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/loglevel"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

func NewLogger(config LoggerConfig) *Logger {
	l := Logger{
		logLevel:      config.LogLevel,
		time:          config.Time,
		outPrefix:     config.OutPrefix,
		outSuffix:     config.OutSuffix,
		debugPrefix:   config.DebugPrefix,
		debugSuffix:   config.DebugSuffix,
		okPrefix:      config.OkPrefix,
		okSuffix:      config.OkSuffix,
		pendingPrefix: config.PendingPrefix,
		pendingSuffix: config.PendingSuffix,
		infoPrefix:    config.InfoPrefix,
		infoSuffix:    config.InfoSuffix,
		warnPrefix:    config.WarnPrefix,
		warnSuffix:    config.WarnSuffix,
		errPrefix:     config.ErrPrefix,
		errSuffix:     config.ErrSuffix,
		fatalPrefix:   config.FatalPrefix,
		fatalSuffix:   config.FatalSuffix,
	}

	if config.Time != nil {
		formatting.TimeLen = max(formatting.TimeLen, len(config.Time.Get()))
	}

	if config.LogLevel == nil {
		l.logLevel = &loglevel.Level
	}

	indentSize := 0
	prefixSize := 0
	indentSize = len(config.Time.Get())

	prefixSize = max(prefixSize, len(config.OutPrefix.Text))
	prefixSize = max(prefixSize, len(config.DebugPrefix.Text))
	prefixSize = max(prefixSize, len(config.OkPrefix.Text))
	prefixSize = max(prefixSize, len(config.PendingPrefix.Text))
	prefixSize = max(prefixSize, len(config.InfoPrefix.Text))
	prefixSize = max(prefixSize, len(config.WarnPrefix.Text))
	prefixSize = max(prefixSize, len(config.ErrPrefix.Text))
	prefixSize = max(prefixSize, len(config.FatalPrefix.Text))

	indentSize += prefixSize
	formatting.IndentSize = max(formatting.IndentSize, indentSize)

	return &l
}

// Prefixes should only contain visible characters as this is used to determine indent size
type LoggerConfig struct {
	LogLevel      *loglevel.LogLevel     // Set log level for this logger. Will default to global log level.
	Time          *timeprefix.TimePrefix // Start logs with current time (UTC or local) with specified format.
	OutPrefix     ansi.ColoredText
	OutSuffix     ansi.ColoredText
	DebugPrefix   ansi.ColoredText
	DebugSuffix   ansi.ColoredText
	OkPrefix      ansi.ColoredText
	OkSuffix      ansi.ColoredText
	PendingPrefix ansi.ColoredText
	PendingSuffix ansi.ColoredText
	InfoPrefix    ansi.ColoredText
	InfoSuffix    ansi.ColoredText
	WarnPrefix    ansi.ColoredText
	WarnSuffix    ansi.ColoredText
	ErrPrefix     ansi.ColoredText
	ErrSuffix     ansi.ColoredText
	FatalPrefix   ansi.ColoredText
	FatalSuffix   ansi.ColoredText
}
