package logger

import (
	"io/fs"
	"log"
	"os"
	cm "persofin/src/commons"
)

/*
global variables
*/

// persofin logger
var persofinLogger *log.Logger

// persofin logger message flag
var logFlag int = log.Ldate | log.Ltime | log.Lshortfile

// persofin log flle mode
var logFileMode fs.FileMode = 0666

//persofin log file pointer
var logFilePtr *os.File

// constant for mode
const (
	INFO    = 0
	DEBUG   = 1
	WARNING = 2
	ERROR   = 3
)

// mode string mapping
var logLevelMap [4]string = [4]string{"INFO", "DEBUG", "WARNING", "ERROR"}

// log level - global variable
var logLevel int = INFO // default is INFO

/*
private functions
*/
// log level - get level
func getLogLevel() int {
	return logLevel
}

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
	flag := os.Getenv(cm.ENABLE_LOGGING_ENV)
	if flag == cm.LOGGING_ENABLED {
		return true
	} else {
		return false
	}
}

/*
public functions
*/

// Public function for Max log level - setup level
func SetupMaxLogLevel(level int) int {
	if level == INFO || level == DEBUG || level == WARNING || level == ERROR {
		logLevel = level
		return cm.SUCCESS
	} else {
		return cm.FAILURE
	}
}

// Public function to print logs
func PrintLog(msgType int, msg string) {
	if getLogLevel() <= msgType && isLoggingEnabled() {
		persofinLogger.Printf("%s:%s", logLevelMap[msgType], msg)
	}
}

// Public function for logging cleanup
func LoggingCleanup() {
	if logFilePtr != nil {
		logFilePtr.Close()
	}
}

// Public function to setup logging
func SetupLogging() int {

	if !isLoggingEnabled() {
		return cm.SUCCESS // exit, as logging is not enabled
	}

	logFileName := readLogFileName() // checking for log file name

	var err error // temp error variable

	logFilePtr, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, logFileMode) // opening file

	// exception handling of log file
	if err != nil {
		log.Fatalf("Issue in reading log file %e", err)
		return cm.FAILURE
	}

	// setting up logger
	persofinLogger = log.New(logFilePtr, "", logFlag)

	// returning
	return cm.SUCCESS
}
