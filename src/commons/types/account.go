package types

import "fmt"

type Account struct {
	Name  string
	Value AccountValue
}

func (av *Account) String() string {
	return fmt.Sprintf("Account { name: %s, currency: %s}", (*av).Name, (*av).Value.String())
}
