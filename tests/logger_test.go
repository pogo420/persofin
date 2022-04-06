package tests

import (
	"io/ioutil"
	"os"
	cm "persofin/src/commons"
	lg "persofin/src/logger"
	"testing"
)

var testLogFile = "psf_test.log"

// Unit testcase for default logging
func TestLoggerDefaultLogFile(t *testing.T) {
	os.Setenv(cm.ENABLE_LOGGING_ENV, cm.LOGGING_ENABLED)
	os.Setenv(cm.LOG_FILE_ENV, testLogFile)
	lg.SetupLogging()
	if lg.SetupLogging() == cm.FAILURE {
		t.Fatalf("Seeing issue in logging")
	}
	lg.LoggingCleanup()
}

// Unit testcase for checking logging not enabled behaviour
func TestLoggingNotEnabled(t *testing.T) {
	os.Setenv(cm.ENABLE_LOGGING_ENV, "")
	lg.SetupLogging()
	if lg.SetupLogging() == cm.FAILURE {
		t.Fatalf("Seeing issue in logging")
	}
	lg.LoggingCleanup()
}

// Unit testcase for log file creation
func TestLoggerDefaultFileWrite(t *testing.T) {
	_ = os.Remove(testLogFile)
	os.Setenv(cm.ENABLE_LOGGING_ENV, cm.LOGGING_ENABLED)
	os.Setenv(cm.LOG_FILE_ENV, testLogFile) // setting log file
	lg.SetupLogging()
	lg.PrintLog(lg.INFO, "TestMessage")
	lg.LoggingCleanup()

	f, err := os.OpenFile(testLogFile, os.O_RDONLY, 0666)
	if err != nil && os.IsNotExist(err) {
		t.Fatalf("Issue in logging into file%e", err)
	} else {
		f.Close()
	}
}

// Unit testcase for logging more check
func TestLoggerLoggingMode(t *testing.T) {
	_ = os.Remove(testLogFile)
	os.Setenv(cm.ENABLE_LOGGING_ENV, cm.LOGGING_ENABLED)
	os.Setenv(cm.LOG_FILE_ENV, testLogFile) // setting log file
	lg.SetupMaxLogLevel(lg.ERROR)
	lg.SetupLogging()
	lg.PrintLog(lg.INFO, "TestMessage")
	lg.LoggingCleanup()

	content, err := ioutil.ReadFile(testLogFile)

	if string(content) != "" || err != nil {
		t.Fatal("Logging levels not working")
	}

}
