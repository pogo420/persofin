package accountsTableReader

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteAccountsTableReader struct {
}

func (s *SqliteAccountsTableReader) AccountsTableRead(account_name string) (*sql.Row, error) {
	return nil, nil
}
