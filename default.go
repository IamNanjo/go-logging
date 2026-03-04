package logging

import (
	"fmt"
	"os"

	"github.com/IamNanjo/go-logging/ansi"
	"github.com/IamNanjo/go-logging/internal/formatting"
	"github.com/IamNanjo/go-logging/loglevel"
)

// Default logger

const (
	outPrefix     = "       "
	debugPrefix   = "       "
	okPrefix      = ansi.Green + "     OK" + ansi.Reset + " "
	pendingPrefix = ansi.Cyan + "PENDING" + ansi.Reset + " "
	infoPrefix    = ansi.Blue + "   INFO" + ansi.Reset + " "
	warnPrefix    = ansi.Yellow + "WARNING" + ansi.Reset + " "
	errPrefix     = ansi.Red + "  ERROR" + ansi.Reset + " "
)

// Writes to stdout regardless of LOGLEVEL. Indents all lines to same level as other default log functions
func Out(format string, args ...any) (int, error) {
	return fmt.Fprintf(os.Stdout, outPrefix+formatting.IndentNewLines(format, len(outPrefix)), args...)
}

// Writes to stdout if LOGLEVEL is at least DEBUG. Indents all lines to same level as other log functions.
func Debug(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.DEBUG {
		return 0, nil
	}
	return fmt.Fprintf(os.Stdout, debugPrefix+formatting.IndentNewLines(format, len(debugPrefix)), args...)
}

// Writes to stdout with green checkmark prefix if LOGLEVEL is at least INFO. Indents all lines to same level.
func Ok(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprintf(os.Stdout, okPrefix+formatting.IndentNewLines(format, len(okPrefix)), args...)
}

// Writes to stdout with green checkmark prefix if LOGLEVEL is at least INFO. Indents all lines to same level.
func Pending(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprintf(os.Stdout, pendingPrefix+formatting.IndentNewLines(format, len(pendingPrefix)), args...)
}

// Writes to stdout with green checkmark prefix if LOGLEVEL is at least INFO. Indents all lines to same level.
func Info(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.INFO {
		return 0, nil
	}
	return fmt.Fprintf(os.Stdout, infoPrefix+formatting.IndentNewLines(format, len(infoPrefix)), args...)
}

// Writes to stderr with yellow warning prefix if LOGLEVEL is at least WARN. Indents all lines to same level.
func Warn(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.WARN {
		return 0, nil
	}
	return fmt.Fprintf(os.Stdout, warnPrefix+formatting.IndentNewLines(format, len(warnPrefix)), args...)
}

// Writes to stderr with red cross prefix if LOGLEVEL is at least ERROR. Indents all lines to same level.
func Err(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.ERROR {
		return 0, nil
	}
	return fmt.Fprintf(os.Stderr, errPrefix+formatting.IndentNewLines(format, len(errPrefix)), args...)
}

// Panic with message. Uses same prefix as Err.
func Fatal(format string, args ...any) {
	panic(fmt.Sprintf(errPrefix+formatting.IndentNewLines(format, len(errPrefix)), args...))
}
