package dbInterface

import (
	interace "persofin/tests/dummy_codes"
)

type InterfaceObject struct {
	Test bool
}

var interfaceObj DbManagerApi = &interace.DummyDbInterface{}  // change it with actual interface
var testInterface DbManagerApi = &interace.DummyDbInterface{} // test mode interface

func (io *InterfaceObject) GetInterfaceObject() DbManagerApi {
	if (*io).Test { // for test mode only
		return testInterface
	}
	return interfaceObj
}
