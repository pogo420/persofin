package tests

import (
	"os"
	am "persofin/src/accountsManager"
	e "persofin/src/commons/exceptions"
	"testing"
)

func TestDbConnectionIssue(t *testing.T) {
	// unit test case to test Db connection issue
	os.Setenv("SQLITE_DB", "")
	_, err := am.GetSqlLiteAccountManager()

	switch ty := err.(type) {
	default:
		t.Fatalf("DB connection error handling not working..")
	case *e.DataBaseConnectionException:
		_ = ty
	}
}

func TestDbConnectionSanity(t *testing.T) {
	// unit test case to test Db connection sanity
	os.Setenv("SQLITE_DB", "infrastructure/test_persofin.db")
	_, err := am.GetSqlLiteAccountManager()

	if err != nil {
		t.Fatalf("DB connection issues")
	}
}
