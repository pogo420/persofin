package tests

import (
	di "persofin/src/dbInterface"
	"testing"
)

func TestDbInterface(t *testing.T) {
	// sanity for getting interface object
	var io *di.InterfaceObject = &di.InterfaceObject{Test: true} //TODO: use it for core

	if (*io).GetInterfaceObject() == nil {
		t.Fatalf("Issue seen in returning interface")
	}
}
