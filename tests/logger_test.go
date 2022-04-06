package tests

import (
	"os"
	cm "persofin/src/commons"
	lg "persofin/src/logger"
	"testing"
)

// unit testcase for default logging
func TestLoggerDefaultLogFile(t *testing.T) {
	os.Setenv(cm.ENABLE_LOGGING_ENV, cm.LOGGING_ENABLED)
	lg.SetupLogging()
	if lg.SetupLogging() == cm.FAILURE {
		t.Fatalf("Seeing issue in logging")
	}
}

// unit testcase for checking logging not enabled behaviour
func TestLoggingNotEnabled(t *testing.T) {
	os.Setenv(cm.ENABLE_LOGGING_ENV, "")
	lg.SetupLogging()
	if lg.SetupLogging() == cm.FAILURE {
		t.Fatalf("Seeing issue in logging")
	}
}
