package tests

import (
	"fmt"
	co "persofin/src/command"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	"testing"
)

// Unit testcase for valid account renaming case
func TestRenameAccountValidCase(t *testing.T) {
	flag := fmt.Sprintf("%s %s", dbi.EXISTS_ACC, dbi.VALID_ACC)
	renameAccoundCommand := &co.RenameAccountCommand{}
	response := renameAccoundCommand.Execute(dbi.GetDbInterface(), flag)

	if response == cmm.FAILURE {
		t.Fatal("Seeing issues in account renaming")
	}
}

// Unit testcase for invalid  account renaming case
func TestRenameAccountInValidCase1(t *testing.T) {
	flag := fmt.Sprintf("%s %s", dbi.VALID_ACC, dbi.EXISTS_ACC)
	renameAccoundCommand := &co.RenameAccountCommand{}
	response := renameAccoundCommand.Execute(dbi.GetDbInterface(), flag)

	if response == cmm.SUCCESS {
		t.Fatal("Seeing issues in account renaming")
	}
}

// Unit testcase for invalid  account renaming case
func TestRenameAccountInValidCase2(t *testing.T) {
	flag := fmt.Sprintf("%s %s", dbi.VALID_ACC, dbi.INVALID_ACC)
	renameAccoundCommand := &co.RenameAccountCommand{}
	response := renameAccoundCommand.Execute(dbi.GetDbInterface(), flag)

	if response == cmm.SUCCESS {
		t.Fatal("Seeing issues in account renaming")
	}
}
