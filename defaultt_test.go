package logging_test

import (
	"testing"

	"github.com/IamNanjo/go-logging"
	"github.com/IamNanjo/go-logging/pkg/format"
	"github.com/IamNanjo/go-logging/pkg/loglevel"
)

var lt = logging.DefaultT()

func TestOutT(t *testing.T) {
	bytesWritten, err := lt.Out("Default output logging \nwith LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestDebugT(t *testing.T) {
	loglevel.Level = loglevel.INFO
	bytesWritten, err := lt.Debug("Default debug logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten != 0 {
		t.Fatalf("Bytes were written regardless of log level")
	}

	loglevel.Level = loglevel.DEBUG
	bytesWritten, err = lt.Debug("Default debug logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestOkT(t *testing.T) {
	bytesWritten, err := lt.Ok("Default ok logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestPendingT(t *testing.T) {
	bytesWritten, err := lt.Pending("Default pending logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestInfoT(t *testing.T) {
	bytesWritten, err := lt.Info("Default info logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestErrT(t *testing.T) {
	err := format.Err("LOGLEVEL set to %s", loglevel.Level)
	err = format.Err("This message uses format.Err() %w", err)
	bytesWritten, err := lt.Err("Default error logging %v\n", err)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatal("No bytes were written")
	}
}

func TestErrWithColonsT(t *testing.T) {
	err := format.Err("Innermost %% error: %.2f", 0.12345)
	err = format.Err("\tError with tab indent %w", err)
	err = format.Err("Outer error:   colon:   and another %w", err)
	err = format.Err("LOGLEVEL set to %s: This message uses format.Err(): It separates errors with colons %w", loglevel.Level, err)
	bytesWritten, err := lt.Err("Default error logging %v\n", err)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatal("No bytes were written")
	}
}

func TestWarnT(t *testing.T) {
	bytesWritten, err := lt.Warn("Default warning logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestFatalT(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("Expected logger to panic but it did not")
		}
	}()

	lt.Fatal("Default fatal logging with LOGLEVEL set to %s\n", loglevel.Level)
}
