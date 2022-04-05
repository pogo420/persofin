package tests

import (
	cp "persofin/src/command_parser"
	"testing"
)

// unit testcase for checking invalid command
func TestCommandProcessor_1(t *testing.T) {
	inp := "dummy"
	if cp.CommandProcessor(inp) != 0 {
		t.Fatalf("Seeing issue with CommandProcessor")
	}
}

// unit testcase for checking exit command
func TestCommandProcessor_2(t *testing.T) {
	inp := "bye"
	if cp.CommandProcessor(inp) != -1 {
		t.Fatalf("Seeing issue with CommandProcessor")
	}
}
