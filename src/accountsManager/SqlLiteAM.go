package accountsManager

import (
	"database/sql"
	"fmt"
	"os"
	e "persofin/src/commons/exceptions"

	_ "github.com/mattn/go-sqlite3"
)

var sqlite_driver_name = "sqlite3"

type SqlLiteAM struct {
	connection *sql.DB
}

func GetSqlLiteAccountManager() (SqlLiteAM, error) {
	// builder function to generate SqlLiteAM struct

	var db_path string = os.Getenv("SQLITE_DB")

	if db_path == "" {
		message := fmt.Sprintf("environmental variable: %s not set", "SQLITE_DB")
		return SqlLiteAM{}, &e.DataBaseConnectionException{Message: message}
	}

	conn, err := sql.Open(sqlite_driver_name, db_path)
	if err != nil {
		message := fmt.Sprintf("Issue in db connection, details: %s", err)
		return SqlLiteAM{}, &e.DataBaseConnectionException{Message: message}
	}
	return SqlLiteAM{connection: conn}, nil
}
