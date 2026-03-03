package logging_test

import (
	"testing"

	"github.com/IamNanjo/logging"
	"github.com/IamNanjo/logging/loglevel"
)

func TestOut(t *testing.T) {
	_, err := logging.Out("Default output logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}

func TestPending(t *testing.T) {
	_, err := logging.Pending("Default pending logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}

func TestInfo(t *testing.T) {
	_, err := logging.Info("Default info logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}

func TestErr(t *testing.T) {
	_, err := logging.Err("Default error logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}

func TestWarn(t *testing.T) {
	_, err := logging.Warn("Default warning logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}

func TestDebug(t *testing.T) {
	_, err := logging.Debug("Default debug logging with LOGLEVEL set to %s\n", loglevel.Level)
	if err != nil {
		t.Fatalf("Logging failed: %v\n", err)
	}
}
