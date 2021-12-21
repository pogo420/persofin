package dbConnApi

import (
	"database/sql"
)

// interface for db connection manager
type DbConnApi interface {
	GetConnection() (*sql.DB, error)
	DeleteExistingConnection()
}
