package types

import "fmt"

type TransactionValue struct {
	Value    int
	Currency Currency
}

func (av *TransactionValue) String() string {
	return fmt.Sprintf("TransactionValue { value: %d, currency: %s}", (*av).Value, (*av).Currency.Value())
}
