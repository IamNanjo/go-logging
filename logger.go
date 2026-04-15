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
	logLevel      loglevel.LogLevel
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
	return l.logLevel
}

// Writes to stdout regardless of LOGLEVEL. Indents all lines to same level as other default log functions
func (l *Logger) Out(format string, args ...any) (int, error) {
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.outPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least DEBUG. Indents all lines to same level as other log functions.
func (l *Logger) Debug(format string, args ...any) (int, error) {
	if l.logLevel < loglevel.DEBUG {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stderr,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.debugPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least INFO. Indents all lines to same level.
func (l *Logger) Ok(format string, args ...any) (int, error) {
	if l.logLevel < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.okPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least INFO. Indents all lines to same level.
func (l *Logger) Pending(format string, args ...any) (int, error) {
	if l.logLevel < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.pendingPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Writes to stdout if LOGLEVEL is at least INFO. Indents all lines to same level.
func (l *Logger) Info(format string, args ...any) (int, error) {
	if l.logLevel < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stdout,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.infoPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Writes to stderr with yellow warning prefix if LOGLEVEL is at least WARN. Indents all lines to same level.
func (l *Logger) Warn(format string, args ...any) (int, error) {
	if l.logLevel < loglevel.WARN {
		return 0, nil
	}
	return fmt.Fprint(
		os.Stderr,
		formatting.IndentWithPrefixAndSuffix(
			l.time,
			l.warnPrefix,
			fmt.Sprintf(format, args...),
			l.errSuffix,
		),
	)
}

// Writes to stderr with red cross prefix if LOGLEVEL is at least ERROR. Indents all lines to same level.
func (l *Logger) Err(format string, args ...any) (int, error) {
	if l.logLevel < loglevel.ERROR {
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
	if l.logLevel < loglevel.CRITICAL {
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
