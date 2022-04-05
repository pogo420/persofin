package db_interface

// Struct for dummy interface
import (
	types "persofin/src/commons"
)

const (
	VALID_ACC   = "VALID_ACC"
	INVALID_ACC = "INVALID_ACC"
	VERSION     = "Dummy Interface v1"
)

type DummyDbInterface struct {
}

func (ddi *DummyDbInterface) GetInterfaceVersion() types.InterfaceVersion {
	return VERSION
}

func (ddi *DummyDbInterface) AccountExists(acc types.AccountName) bool {
	if acc == VALID_ACC {
		return false
	} else {
		return true
	}
}

func (ddi *DummyDbInterface) CreateAccount(acc types.AccountName) int {
	if acc == VALID_ACC {
		return 0
	} else {
		return -1
	}
}

func (ddi *DummyDbInterface) GetAccountBalance(acc types.AccountName) types.TansactionValue {
	if acc == VALID_ACC {
		return 100
	} else {
		return -100
	}
}

func (ddi *DummyDbInterface) GetAllAccountsBalance() types.AccountStats {
	return map[types.AccountName]types.TansactionValue{
		VALID_ACC:        34,
		VALID_ACC + "_1": 56,
	}
}

func (ddi *DummyDbInterface) RenameAccount(acc1 types.AccountName, acc2 types.AccountName) int {
	if acc1 != INVALID_ACC && acc2 != INVALID_ACC {
		return 0
	} else {
		return -1
	}
}
func (ddi *DummyDbInterface) AddTransaction(d types.Date, acc types.AccountName, val types.TansactionValue) int {
	if acc == VALID_ACC {
		return 0
	} else {
		return -1
	}
}

func (ddi *DummyDbInterface) AccountTransaction(d types.Date, acc1 types.AccountName, acc2 types.AccountName, val types.TansactionValue) int {
	if acc1 != INVALID_ACC && acc2 != INVALID_ACC {
		return 0
	} else {
		return -1
	}
}
