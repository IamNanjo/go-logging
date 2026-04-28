package logging_test

import (
	"testing"

	"github.com/IamNanjo/go-logging"
	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/format"
	"github.com/IamNanjo/go-logging/pkg/loglevel"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

var cl *logging.Logger = logging.NewLogger(logging.LoggerConfig{
	LogLevel: new(loglevel.DEBUG),
	Time:     &timeprefix.TimePrefix{UTC: true, Format: "2006-01-02 15:04"},

	OutPrefix:     ansi.ColoredText{Color: ansi.None, Text: "Out prefix"},
	OutSuffix:     ansi.ColoredText{Color: ansi.None, Text: " [Out suffix]"},
	DebugPrefix:   ansi.ColoredText{Color: ansi.Purple, Text: "Debug prefix"},
	DebugSuffix:   ansi.ColoredText{Color: ansi.Purple, Text: " [Debug suffix]"},
	OkPrefix:      ansi.ColoredText{Color: ansi.Green, Text: "Ok prefix"},
	OkSuffix:      ansi.ColoredText{Color: ansi.Green, Text: " [Ok suffix]"},
	PendingPrefix: ansi.ColoredText{Color: ansi.Cyan, Text: "Pending prefix"},
	PendingSuffix: ansi.ColoredText{Color: ansi.Cyan, Text: " [Pending suffix]"},
	InfoPrefix:    ansi.ColoredText{Color: ansi.Blue, Text: "Info prefix"},
	InfoSuffix:    ansi.ColoredText{Color: ansi.Blue, Text: " [Info suffix]"},
	WarnPrefix:    ansi.ColoredText{Color: ansi.Yellow, Text: "Warn prefix"},
	WarnSuffix:    ansi.ColoredText{Color: ansi.Yellow, Text: " [Warn suffix]"},
	ErrPrefix:     ansi.ColoredText{Color: ansi.Red, Text: "Err prefix"},
	ErrSuffix:     ansi.ColoredText{Color: ansi.Red, Text: " [Err suffix]"},
	FatalPrefix:   ansi.ColoredText{Color: ansi.Red, Text: "Fatal prefix"},
	FatalSuffix:   ansi.ColoredText{Color: ansi.Red, Text: " [Fatal suffix]"},
})

func TestNewOut(t *testing.T) {
	bytesWritten, err := cl.Out("Custom output logging with LOGLEVEL set to %s\n", cl.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Out")
	}
}

func TestNewDebug(t *testing.T) {
	bytesWritten, err := cl.Debug("Custom debug logging with LOGLEVEL set to %s\n", cl.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Debug")
	}
}
func TestNewOk(t *testing.T) {
	bytesWritten, err := cl.Ok("Custom ok logging\nwith LOGLEVEL set to %s\n", cl.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Ok")
	}
}
func TestNewInfo(t *testing.T) {
	bytesWritten, err := cl.Info("Custom info logging with LOGLEVEL set to %s\n", cl.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Info")
	}
}
func TestNewPending(t *testing.T) {
	bytesWritten, err := cl.Pending("Custom pending logging with LOGLEVEL set to %s\n", cl.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Pending")
	}
}
func TestNewWarn(t *testing.T) {
	bytesWritten, err := cl.Warn("Custom warning logging with LOGLEVEL set to %s\n", cl.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Warn")
	}
}
func TestNewErr(t *testing.T) {
	err := format.Err("LOGLEVEL set to %s", cl.LogLevel())
	err = format.Err("This message uses format.Err() %w", err)
	bytesWritten, err := cl.Err("Default error logging %v\n", err)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatal("No bytes were written with Logger.Err")
	}
}
func TestNewFatal(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("Expected Logger.Fatal to panic but it did not")
		}
	}()

	cl.Fatal("Custom fatal logging")
}
