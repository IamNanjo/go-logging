package logging

import (
	"fmt"
	"os"

	"github.com/IamNanjo/go-logging/internal/formatting"
	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/loglevel"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

type Logger struct {
	logLevel      *loglevel.LogLevel
	time          *timeprefix.TimePrefix
	outPrefix     ansi.ColoredText
	outSuffix     ansi.ColoredText
	debugPrefix   ansi.ColoredText
	debugSuffix   ansi.ColoredText
	okPrefix      ansi.ColoredText
	okSuffix      ansi.ColoredText
	pendingPrefix ansi.ColoredText
	pendingSuffix ansi.ColoredText
	infoPrefix    ansi.ColoredText
	infoSuffix    ansi.ColoredText
	warnPrefix    ansi.ColoredText
	warnSuffix    ansi.ColoredText
	errPrefix     ansi.ColoredText
	errSuffix     ansi.ColoredText
	fatalPrefix   ansi.ColoredText
	fatalSuffix   ansi.ColoredText
}

func (l *Logger) LogLevel() loglevel.LogLevel {
	if l == nil || l.logLevel == nil {
		return loglevel.Level
	}
	return *l.logLevel
}

// Return config the config for this logger.
// Pointer values are copied so they cannot be used to edit existing loggers.
func (l *Logger) Config() *LoggerConfig {
	if l == nil {
		return nil
	}

	var ll = l.logLevel
	if ll != &loglevel.Level {
		lll := *ll
		ll = &lll
	}

	var t *timeprefix.TimePrefix
	if l.time != nil {
		lt := *l.time
		t = &lt
	}

	return &LoggerConfig{
		LogLevel:      ll,
		Time:          t,
		OutPrefix:     l.outPrefix,
		OutSuffix:     l.outSuffix,
		DebugPrefix:   l.debugPrefix,
		DebugSuffix:   l.debugSuffix,
		OkPrefix:      l.okPrefix,
		OkSuffix:      l.okSuffix,
		PendingPrefix: l.pendingPrefix,
		PendingSuffix: l.pendingSuffix,
		InfoPrefix:    l.infoPrefix,
		InfoSuffix:    l.infoSuffix,
		WarnPrefix:    l.warnPrefix,
		WarnSuffix:    l.warnSuffix,
		ErrPrefix:     l.errPrefix,
		ErrSuffix:     l.errSuffix,
		FatalPrefix:   l.fatalPrefix,
		FatalSuffix:   l.fatalSuffix,
	}
}

// Writes to stdout regardless of LOGLEVEL. Indents all lines to same level as other default log functions
func (l *Logger) Out(format string, args ...any) (int, error) {
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.outPrefix,
			fmt.Sprintf(format, args...),
			l.outSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least DEBUG. Indents all lines to same level as other log functions.
func (l *Logger) Debug(format string, args ...any) (int, error) {
	if *l.logLevel < loglevel.DEBUG {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stderr,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.debugPrefix,
			fmt.Sprintf(format, args...),
			l.debugSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least INFO. Indents all lines to same level.
func (l *Logger) Ok(format string, args ...any) (int, error) {
	if *l.logLevel < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.okPrefix,
			fmt.Sprintf(format, args...),
			l.okSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least INFO. Indents all lines to same level.
func (l *Logger) Pending(format string, args ...any) (int, error) {
	if *l.logLevel < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.pendingPrefix,
			fmt.Sprintf(format, args...),
			l.pendingSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least INFO. Indents all lines to same level.
func (l *Logger) Info(format string, args ...any) (int, error) {
	if *l.logLevel < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.infoPrefix,
			fmt.Sprintf(format, args...),
			l.infoSuffix,
		),
	)
}

// Writes to stderr if LOGLEVEL is at least WARN. Indents all lines to same level.
func (l *Logger) Warn(format string, args ...any) (int, error) {
	if *l.logLevel < loglevel.WARN {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stderr,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.warnPrefix,
			fmt.Sprintf(format, args...),
			l.warnSuffix,
		),
	)
}

// Writes to stderr if LOGLEVEL is at least ERROR. Indents all lines to same level.
func (l *Logger) Err(format string, args ...any) (int, error) {
	if *l.logLevel < loglevel.ERROR {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stderr,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.errPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Same as Err but with panic.
func (l *Logger) Fatal(format string, args ...any) {
	if *l.logLevel < loglevel.CRITICAL {
		panic("")
	}
	panic(
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.errPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}
