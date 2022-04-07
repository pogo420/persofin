package tests

import (
	"os"
	cm "persofin/src/command"
	com "persofin/src/commons"
	dbi "persofin/src/db_interface"
	"testing"
)

// unit testcase for AccountCreation Success
func TestAccountCreationSuccess(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.CreateAccountCommand{}
	response := command.Execute(dbi.VALID_ACC)
	if response == com.FAILURE {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for AccountCreation Failure - invalid account
func TestAccountCreationFailureInvalAcc(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.CreateAccountCommand{}
	response := command.Execute(dbi.INVALID_ACC)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for AccountCreation Failure - account exists
func TestAccountCreationFailureAccExists(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.CreateAccountCommand{}
	response := command.Execute(dbi.EXISTS_ACC)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}
