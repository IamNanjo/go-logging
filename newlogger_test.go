package logging_test

import (
	"testing"
	"time"

	"github.com/IamNanjo/go-logging"
	"github.com/IamNanjo/go-logging/pkg/ansi"
	"github.com/IamNanjo/go-logging/pkg/format"
	"github.com/IamNanjo/go-logging/pkg/loglevel"
	"github.com/IamNanjo/go-logging/pkg/timeprefix"
)

func TestNew(t *testing.T) {
	l := logging.NewLogger(logging.LoggerConfig{
		LogLevel: loglevel.DEBUG,
		Time:     &timeprefix.TimePrefix{UTC: false, Format: time.RFC3339},

		OutPrefix:     ansi.ColoredText{Color: ansi.None, Text: "Out prefix "},
		OutSuffix:     ansi.ColoredText{Color: ansi.None, Text: " Out suffix"},
		DebugPrefix:   ansi.ColoredText{Color: ansi.Purple, Text: "Debug prefix "},
		DebugSuffix:   ansi.ColoredText{Color: ansi.Purple, Text: " Debug suffix"},
		OkPrefix:      ansi.ColoredText{Color: ansi.Green, Text: "Ok prefix "},
		OkSuffix:      ansi.ColoredText{Color: ansi.Green, Text: " Ok suffix"},
		PendingPrefix: ansi.ColoredText{Color: ansi.Cyan, Text: "Pending prefix "},
		PendingSuffix: ansi.ColoredText{Color: ansi.Cyan, Text: " Pending suffix"},
		InfoPrefix:    ansi.ColoredText{Color: ansi.Blue, Text: "Info prefix "},
		InfoSuffix:    ansi.ColoredText{Color: ansi.Blue, Text: " Info suffix"},
		WarnPrefix:    ansi.ColoredText{Color: ansi.Yellow, Text: "Warn prefix "},
		WarnSuffix:    ansi.ColoredText{Color: ansi.Yellow, Text: " Warn suffix"},
		ErrPrefix:     ansi.ColoredText{Color: ansi.Red, Text: "Err prefix "},
		ErrSuffix:     ansi.ColoredText{Color: ansi.Red, Text: " Err suffix"},
		FatalPrefix:   ansi.ColoredText{Color: ansi.Red, Text: "Fatal prefix "},
		FatalSuffix:   ansi.ColoredText{Color: ansi.Red, Text: " Fatal suffix"},
	})

	bytesWritten, err := l.Out("Custom output logging with LOGLEVEL set to %s\n", l.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}

	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Out")
	}

	bytesWritten, err = l.Debug("Custom debug logging with LOGLEVEL set to %s\n", l.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Debug")
	}

	bytesWritten, err = l.Ok("Custom ok logging\nwith LOGLEVEL set to %s\n", l.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Ok")
	}

	bytesWritten, err = l.Info("Custom info logging with LOGLEVEL set to %s\n", l.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Info")
	}

	bytesWritten, err = l.Pending("Custom pending logging with LOGLEVEL set to %s\n", l.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Pending")
	}

	bytesWritten, err = l.Warn("Custom warning logging with LOGLEVEL set to %s\n", l.LogLevel())
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written with Logger.Warn")
	}

	err = format.Err("LOGLEVEL set to %s", l.LogLevel())
	err = format.Err("This message uses format.Err() %w", err)
	bytesWritten, err = l.Err("Default error logging %v\n", err)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatal("No bytes were written with Logger.Err")
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Expected Logger.Fatal to panic but it did not")
		}
	}()

	l.Fatal("Custom fatal logging")
}
