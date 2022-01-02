package dbManagerApi

import ty "persofin/src/commons/types"

type DbManagerApi interface {
	ReadAccountsTable(string) *ty.Account
	WriteAccountsTable(*ty.Account) int
	ReadTransactionTableId(int) *ty.Transaction
	ReadTransactionTableByDate(ty.Ts, ty.Ts) *([]ty.Transaction)
	WriteTransactionTable(*ty.Transaction) int
}
