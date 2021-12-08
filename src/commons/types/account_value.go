package types

import "fmt"

type AccountValue struct {
	Value    int
	Currency Currency
}

func (av *AccountValue) String() string {
	return fmt.Sprintf("AccountValue { value: %d, currency: %s}", (*av).Value, (*av).Currency.Value())
}
