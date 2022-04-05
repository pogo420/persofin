package db_interface

import (
	types "persofin/src/commons"
)

// Base interface all DB interfaces must implement it.
type BaseDbInterface interface {
	// Getting interface version
	GetInterfaceVersion() types.InterfaceVersion
	// Is account exists
	AccountExists(types.AccountName) bool
	// Creating accoung
	CreateAccount(types.AccountName) int // 0 - success , -1 failure
	// Get a account balance
	GetAccountBalance(types.AccountName) types.TansactionValue
	// Get all account balance
	GetAllAccountsBalance() types.AccountStats
	// Rename account
	RenameAccount(types.AccountName, types.AccountName) int // 0 - success , -1 failure
	// Add transaction
	AddTransaction(types.Date, types.AccountName, types.TansactionValue) int // 0 - success , -1 failure
	// Account transfer
	AccountTransaction(types.Date, types.AccountName, types.AccountName, types.TansactionValue) int // 0 - success , -1 failure
}
