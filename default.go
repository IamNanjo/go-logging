package logging

import (
	"fmt"
	"os"

	"github.com/IamNanjo/go-logging/ansi"
	"github.com/IamNanjo/go-logging/icons"
	"github.com/IamNanjo/go-logging/internal/formatting"
	"github.com/IamNanjo/go-logging/loglevel"
)

// Default logger

const outPrefix = "   "
const okPrefix = ansi.Green + icons.Ok + ansi.Reset + " "
const pendingPrefix = ansi.Cyan + icons.Pending + ansi.Reset + " "
const infoPrefix = ansi.Blue + icons.Info + ansi.Reset + " "
const warnPrefix = ansi.Yellow + icons.Warn + ansi.Reset + " "
const errPrefix = ansi.Red + icons.Err + ansi.Reset + " "

// Writes to stdout regardless of LOGLEVEL. Indents all lines to same level as other default log functions
func Out(format string, args ...any) (int, error) {
	return fmt.Fprintf(os.Stdout, outPrefix+formatting.IndentNewLines(format, len(outPrefix)), args...)
}

// Writes to stdout if LOGLEVEL is at least DEBUG. Indents all lines to same level as other log functions with no icon.
func Debug(format string, args ...any) (int, error) {
	if loglevel.Level < loglevel.DEBUG {
		return 0, nil
	}
	return fmt.Fprintf(os.Stdout, outPrefix+formatting.IndentNewLines(format, len(outPrefix)), args...)
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

// Same as Err but exits program with code 1 and outputs regardless of LOGLEVEL (CRITICAL is the minimum value).
func Fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, errPrefix+formatting.IndentNewLines(format, len(errPrefix)), args...)
	os.Exit(1)
}
