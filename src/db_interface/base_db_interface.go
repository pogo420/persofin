package db_interface

import (
	types "persofin/src/commons"
)

// Base interface all DB interfaces must implement it.
type BaseDbInterface interface {
	AccountExists(types.AccountName) bool
	CreateAccount(types.AccountName) int // 0 - success , -1 failure
	GetAccountBalance(types.AccountName) types.TansactionValue
	GetAllAccountsBalance() types.AccountStats
	RenameAccount(types.AccountName, types.AccountName) int                                         // 0 - success , -1 failure
	AddTransaction(types.Date, types.AccountName, types.TansactionValue) int                        // 0 - success , -1 failure
	AccountTransaction(types.Date, types.AccountName, types.AccountName, types.TansactionValue) int // 0 - success , -1 failure
}
