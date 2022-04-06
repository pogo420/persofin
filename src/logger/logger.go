package logger

import (
	"log"
	"os"
	cm "persofin/src/commons"
)

// function to read log file name
func readLogFileName() string {
	log_file := os.Getenv(cm.LOG_FILE_ENV)
	if log_file != "" {
		return log_file
	} else {
		return cm.DEFAULT_LOG_FILE
	}
}

// function to check logging is enabled or not
func isLoggingEnabled() bool {
	flag := os.Getenv(cm.LOGGING_ENABLED)
	if flag == cm.LOGGING_ENABLED {
		return true
	} else {
		return false
	}
}

// Public function to setup logging
func SetupLogging() int {

	if !isLoggingEnabled() {
		return cm.SUCCESS // exit, as logging is not enabled
	}

	log_file := readLogFileName() // checking for log file name

	file, err := os.OpenFile(log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // opening file

	if err != nil {
		log.Fatalf("Issue in reading log file %e", err)
		return cm.FAILURE
	}

	log.SetOutput(file) // setting log file
	return cm.SUCCESS
}
