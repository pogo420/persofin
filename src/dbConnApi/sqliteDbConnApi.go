package dbConnApi

import (
	"database/sql"
	"fmt"
	e "persofin/src/commons/exceptions"
	ut "persofin/src/commons/utils"

	_ "github.com/mattn/go-sqlite3"
)

//variable to keep track of connection
var connection *sql.DB

// DB specific constants
const Sqlite_driver_env string = "SQLITE_DRIVER_NAME"
const Sqlite_db_env string = "SQLITE_DB"

type SqliteDbConnApi struct {
}

func (c *SqliteDbConnApi) DeleteExistingConnection() {
	connection = nil
}

func (c *SqliteDbConnApi) GetConnection() (*sql.DB, error) {
	if connection != nil {
		return connection, nil
	} else {

		// driver name
		driver_name := ut.ReadEnvVariable(Sqlite_driver_env)
		if driver_name == "" {
			message := fmt.Sprintf("environmental variable: %s not set", Sqlite_driver_env)
			return nil, &e.DataBaseConnectionException{Message: message}
		}

		// database name
		database_name := ut.ReadEnvVariable(Sqlite_db_env)
		if database_name == "" {
			message := fmt.Sprintf("environmental variable: %s not set", Sqlite_db_env)
			return nil, &e.DataBaseConnectionException{Message: message}
		}

		// connection generation
		var err error
		connection, err = sql.Open(driver_name, database_name)

		if err != nil {
			message := fmt.Sprintf("Issue in database connection: %s", err)
			return nil, &e.DataBaseConnectionException{Message: message}
		}

		return connection, nil
	}
}
