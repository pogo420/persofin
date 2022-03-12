package tests

import (
	"os"
	ut "persofin/src/commons/utils"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// test case to env variable read
	os.Setenv("A", "Ola")
	if ut.ReadEnvVariable("A") != "Ola" {
		t.Fatalf("Not able to read env variable..!")
	}
}

func TestGetEnvIssue(t *testing.T) {
	// test case to env variable wrong read
	os.Setenv("A", "Ola")
	if ut.ReadEnvVariable("B") != "" {
		t.Fatalf("Not able to read env variable..!")
	}
}

func TestModeVariablePos(t *testing.T) {
	// positive testing of test mode checker
	os.Setenv(ut.Test_mode_flag, "0")
	if ut.IsTestMode() {
		t.Fatalf("Issue in test mode checker")
	}
}

func TestModeVariableNeg(t *testing.T) {
	// negative testing of test mode checker
	os.Setenv(ut.Test_mode_flag, "1")
	if !ut.IsTestMode() {
		t.Fatalf("Issue in test mode checker")
	}
}
