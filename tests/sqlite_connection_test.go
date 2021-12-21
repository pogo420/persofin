package tests

import (
	"os"
	conn "persofin/src/dbConnApi"
	"testing"
)

func TestSqliteDbConnectionSanity(t *testing.T) {
	// unit test case to test Db connection sanity
	os.Setenv("SQLITE_DB", "infrastructure/test_persofin.db")
	os.Setenv("SQLITE_DRIVER_NAME", "sqlite3")

	var db_conn conn.DbConnApi = &conn.SqliteDbConnApi{} // pointer receiver method

	_, err := db_conn.GetConnection()

	if err != nil {
		t.Fatalf("DB connection issues %s", err)
	}
}

func TestSqliteDbConnectionExceptionCheck1(t *testing.T) {
	// unit test case to test Db connection sanity
	os.Setenv("SQLITE_DB", "")
	os.Setenv("SQLITE_DRIVER_NAME", "sqlite3")

	var db_conn conn.DbConnApi = &conn.SqliteDbConnApi{} // pointer receiver method

	db_conn.DeleteExistingConnection() // deleting existing connection
	_, err := db_conn.GetConnection()

	if err == nil {
		t.Fatal("Connection API error handling not working")
	}
}

func TestSqliteDbConnectionExceptionCheck2(t *testing.T) {
	// unit test case to test Db connection sanity
	os.Setenv("SQLITE_DB", "infrastructure/test_persofin.db")
	os.Setenv("SQLITE_DRIVER_NAME", "")

	var db_conn conn.DbConnApi = &conn.SqliteDbConnApi{} // pointer receiver method

	db_conn.DeleteExistingConnection() // deleting existing connection
	_, err := db_conn.GetConnection()

	if err == nil {
		t.Fatal("Connection API error handling not working")
	}
}
