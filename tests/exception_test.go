package tests

import (
	e "persofin/src/commons/exceptions"
	"testing"
)

func exception_driver() []error {
	var erros = []error{
		&e.InvalidCommandExceptions{},
		&e.TransactionNotFoundException{},
		&e.TransactionCreationException{},
		&e.AccountCreationException{},
		&e.AccountNotFoundException{},
		&e.AccountUpdateException{},
		&e.InsufficientBalanceException{},
		&e.InvalidCommandExceptions{},
		&e.DataBaseConnectionException{},
	}
	return erros
}

func TestException(t *testing.T) {
	errors := exception_driver()

	for _, e := range errors {
		if e == nil {
			t.Fatalf("Not raising error %s", e)
		}
	}
}
