package accountsTableReader

import (
	"database/sql"
)

// Interface for Accounts Table reading
type AccountsTableReader interface {
	AccountsTableRead(account_name string) (*sql.Row, error)
}
