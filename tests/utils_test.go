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
