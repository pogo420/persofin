package accountsManager

import ty "persofin/src/commons/types"

type AccountsManager interface {
	// interface for all account manager logic
	createAccount(name string) (ty.Response, error)
	getBalance(name string) (ty.AccountValue, error)
	updateBalance(name string, av ty.AccountValue) (ty.Response, error)
	getAccount(name string) (ty.Account, error)
	renameAccount(oldName string, newName string) (ty.Response, error)
}
