package tests

import (
	"fmt"
	"os"
	cm "persofin/src/command"
	com "persofin/src/commons"
	dbi "persofin/src/db_interface"
	"testing"
)

// // unit testcase for checking invalid date
// func TestAddTransInvaidDate(t *testing.T) {
// 	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
// 	command := &cm.AddTransactionCommand{}

// 	flag := fmt.Sprintf("%s %s %s", "439-234-89", dbi.VALID_ACC, "4456")
// 	response := command.Execute(dbi.GetDbInterface(), flag)
// 	if response == com.SUCCESS {
// 		t.Fatal("Issue in create account command")
// 	}
// }

// // unit testcase for checking invalid flags
// func TestAddTransInvaidFlags(t *testing.T) {
// 	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
// 	command := &cm.AddTransactionCommand{}

// 	flag := fmt.Sprintf("%s %s-%s", "439-234-89", dbi.VALID_ACC, "4456")
// 	response := command.Execute(dbi.GetDbInterface(), flag)
// 	if response == com.SUCCESS {
// 		t.Fatal("Issue in create account command")
// 	}
// }

// // unit testcase for checking response of invalid account
// func TestAddTransInvalidAccount(t *testing.T) {
// 	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
// 	command := &cm.AddTransactionCommand{}

// 	flag := fmt.Sprintf("%s %s %s", "2022-04-12", dbi.INVALID_ACC, "4456")
// 	response := command.Execute(dbi.GetDbInterface(), flag)
// 	if response == com.SUCCESS {
// 		t.Fatal("Issue in create account command")
// 	}
// }

// // unit testcase for checking response of invalid tarnsaction
// func TestAddTransInvalidTransaction(t *testing.T) {
// 	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
// 	command := &cm.AddTransactionCommand{}

// 	flag := fmt.Sprintf("%s %s %s", "2022-04-12", dbi.VALID_ACC, "445-689")
// 	response := command.Execute(dbi.GetDbInterface(), flag)
// 	if response == com.SUCCESS {
// 		t.Fatal("Issue in create account command")
// 	}
// }

// unit testcase for checking response of valid tarnsaction
func TestAddTransValidTransaction(t *testing.T) {
	os.Setenv(com.TEST_MODE_ENV, com.TEST_MODE_ENABLED)
	command := &cm.AddTransactionCommand{}

	flag := fmt.Sprintf("%s %s %s", "2022-04-12", dbi.EXISTS_ACC, "4456")
	fmt.Println(flag)
	response := command.Execute(dbi.GetDbInterface(), flag)
	if response == com.FAILURE {
		t.Fatal("Issue in create account command")
	}
}
