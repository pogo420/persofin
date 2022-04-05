package db_interface

import (
	"os"
	cm "persofin/src/commons"
)

var TEST_INTERFACE BaseDbInterface = &DummyDbInterface{}
var PROD_INTERFACE BaseDbInterface = &DummyDbInterface{}

// function to check test mode
func getTestFlag() bool {
	if os.Getenv(cm.TEST_MODE_ENV) == "1" {
		return true
	} else {
		return false
	}
}

// Public function to return interface object
func GetDbInterface() BaseDbInterface {
	if getTestFlag() {
		return TEST_INTERFACE
	}
	return PROD_INTERFACE
}
