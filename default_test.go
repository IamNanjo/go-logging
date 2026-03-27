package logging_test

import (
	"testing"

	"github.com/IamNanjo/go-logging"
	"github.com/IamNanjo/go-logging/pkg/loglevel"
)

func TestOut(t *testing.T) {
	_, err := logging.Out("Default output logging \nwith LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}

func TestDebug(t *testing.T) {
	_, err := logging.Debug("Default debug logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
}

func TestOk(t *testing.T) {
	_, err := logging.Ok("Default ok logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
}

func TestPending(t *testing.T) {
	bytesWritten, err := logging.Pending("Default pending logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
	if bytesWritten == 0 {
		t.Fatalf("No bytes were written")
	}
}

func TestInfo(t *testing.T) {
	_, err := logging.Info("Default info logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
}

func TestErr(t *testing.T) {
	_, err := logging.Err("Default error logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
}

func TestWarn(t *testing.T) {
	_, err := logging.Warn("Default warning logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v", err)
	}
}
