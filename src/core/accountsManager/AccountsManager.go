package accountsManager

import ty "persofin/src/commons/types"

type AccountsManager interface {
	// interface for all account manager logic
	CreateAccount(name string) (ty.Response, error)
	GetBalance(name string) (ty.AccountValue, error)
	UpdateBalance(name string, av ty.AccountValue) (ty.Response, error)
	GetAccount(name string) (ty.Account, error)
	RenameAccount(oldName string, newName string) (ty.Response, error)
}
