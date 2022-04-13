package tests

import (
	"fmt"
	"os"
	cm "persofin/src/command"
	com "persofin/src/commons"
	dbi "persofin/src/db_interface"
	"testing"
)

// unit testcase for checking invalid date
func TestAccTransInvaidDate(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AccountTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s %s", "9088-94-12", dbi.EXISTS_ACC, dbi.EXISTS_ACC, "4456")
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for checking invalid Account
func TestAccTransInvaidAccount(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AccountTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s %s", "2022-04-12", dbi.INVALID_ACC, dbi.EXISTS_ACC, "4456")
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for checking invalid transaction
func TestAccTransInvaidTransaction(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AccountTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s %s", "2022-04-12", dbi.EXISTS_ACC, dbi.EXISTS_ACC, "44-56")
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for checking invalid same account
func TestAccTransInvaidSameAccount(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AccountTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s %s", "2022-04-12", dbi.EXISTS_ACC, dbi.EXISTS_ACC, "4456")
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for checking invalid flag
func TestAccTransInvaidFlag(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AccountTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s%s", "2022-04-12", dbi.EXISTS_ACC, dbi.EXISTS_ACC, "4456")
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.SUCCESS {
		t.Fatal("Issue in create account command")
	}
}

// unit testcase for checking valid flow
func TestAccTransValidFlow(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AccountTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s %s", "2022-04-12", dbi.EXISTS_ACC, dbi.EXISTS_ACC_2, "4456")
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.FAILURE {
		t.Fatal("Issue in create account command")
	}
}
