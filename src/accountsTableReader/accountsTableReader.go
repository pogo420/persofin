package accountsTableReader

import ty "persofin/src/commons/types"

// Interface for Accounts Table reading
type AccountsTableReader interface {
	AccountsTableRead(account_name string) (*ty.Account, error)
}
