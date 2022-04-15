package tests

import (
	"os"
	cm "persofin/src/command"
	com "persofin/src/commons"
	dbi "persofin/src/db_interface"
	"testing"
)

// unit test for get balance success
func TestGetBalanceSuccess(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.GetBalanceCommand{}
	response := command.Execute(dbi.GetDbInterface(), dbi.EXISTS_ACC)
	if response == com.FAILURE {
		t.Fatal("Issue in getting account balance")
	}
}

// unit test for get balance failure - Invalid account name
func TestGetBalanceFailure(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.GetBalanceCommand{}
	response := command.Execute(dbi.GetDbInterface(), dbi.INVALID_ACC)
	if response == com.SUCCESS {
		t.Fatal("Issue in getting account balance")
	}
}
