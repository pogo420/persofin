package exceptions

import "fmt"

type InvalidCommandExceptions struct {
	message string
}

func (e *InvalidCommandExceptions) Error() string {
	return fmt.Sprintf("Invalid commad used: %s", (*e).message)
}

type TransactionNotFoundException struct {
	message string
}

func (e *TransactionNotFoundException) Error() string {
	return fmt.Sprintf("Not able to find the transaction specified: %s", (*e).message)
}

type TransactionCreationException struct {
	message string
}

func (e *TransactionCreationException) Error() string {
	return fmt.Sprintf("Not able to create the transaction: %s", (*e).message)
}

type InsufficientBalanceException struct {
	message string
}

func (e *InsufficientBalanceException) Error() string {
	return fmt.Sprintf("Insufficient Balance: %s", (*e).message)
}

type AccountCreationException struct {
	message string
}

func (e *AccountCreationException) Error() string {
	return fmt.Sprintf("Not able to create account: %s", (*e).message)
}

type AccountNotFoundException struct {
	message string
}

func (e *AccountNotFoundException) Error() string {
	return fmt.Sprintf("Not able to find the account: %s", (*e).message)
}

type AccountUpdateException struct {
	message string
}

func (e *AccountUpdateException) Error() string {
	return fmt.Sprintf("Not able to update the account: %s", (*e).message)
}
