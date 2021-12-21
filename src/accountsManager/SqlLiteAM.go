package accountsManager

import (
	"database/sql"
	"fmt"
	c "persofin/src/commons/constants"
	e "persofin/src/commons/exceptions"
	ty "persofin/src/commons/types"

	_ "github.com/mattn/go-sqlite3"
)

type SqlLiteAM struct {
	connection *sql.DB
}

// func GetSqlLiteAccountManager() (SqlLiteAM, error) {
// 	// builder function to generate SqlLiteAM struct

// 	var db_path string = os.Getenv(c.Sqlite_db_env)

// 	if db_path == "" {
// 		message := fmt.Sprintf("environmental variable: %s not set", c.Sqlite_db_env)
// 		return SqlLiteAM{}, &e.DataBaseConnectionException{Message: message}
// 	}

// 	conn, err := sql.Open(c.Sqlite_driver_name, db_path)
// 	if err != nil {
// 		message := fmt.Sprintf("Issue in db connection, details: %s", err)
// 		return SqlLiteAM{}, &e.DataBaseConnectionException{Message: message}
// 	}
// 	return SqlLiteAM{connection: conn}, nil
// }

func (sam *SqlLiteAM) GetAccount(name string) (ty.Account, error) {
	// method to get account information

	account_query := fmt.Sprintf("select * from %s where %s=%s", c.ACCOUNTS_TABLE, c.ACCOUNTS_NAME_COLUMN, name)
	rows, err := sam.connection.Query(account_query)

	if err != nil || rows == nil {
		return ty.Account{}, &e.AccountNotFoundException{Message: "Issue in account name"}
	}

	var account_name string
	var timestamp string
	var account_value int

	for rows.Next() {
		err = rows.Scan(&account_name, &timestamp, account_value)
		break
	}
	if err != nil {
		return ty.Account{}, &e.AccountNotFoundException{Message: "Issue in account name"}
	}

	return ty.Account{Name: account_name, Value: ty.AccountValue{Value: account_value, Currency: ty.INR}}, nil

}
