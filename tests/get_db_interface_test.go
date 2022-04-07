package tests

import (
	"os"
	cm "persofin/src/commons"
	di "persofin/src/db_interface"
	"testing"
)

// unit testcase for checking get interface function
func TestGetInterface(t *testing.T) {
	os.Setenv(cm.TEST_MODE_ENV, cm.TEST_MODE_ENABLED)
	version := "Dummy Interface v1"
	if di.GetDbInterface().GetInterfaceVersion() != cm.InterfaceVersion(version) {
		t.Fatal("Issue in GetInterface function!")
	}
}
