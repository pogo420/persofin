package dbInterface

import (
	ut "persofin/src/commons/utils"
	interace "persofin/tests/dummy_codes"
)

type InterfaceObject struct {
}

var interfaceObj DbManagerApi = &interace.DummyDbInterface{}  // change it with actual interface
var testInterface DbManagerApi = &interace.DummyDbInterface{} // test mode interface

func (io *InterfaceObject) GetInterfaceObject() DbManagerApi {
	if ut.IsTestMode() { // for test mode only
		return testInterface
	} else {
		return interfaceObj
	}
}
