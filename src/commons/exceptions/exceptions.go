package exceptions

import "fmt"

type InvalidCommandExceptions struct {
	Message string
}

func (e *InvalidCommandExceptions) Error() string {
	return fmt.Sprintf("Invalid commad used: %s", (*e).Message)
}

type TransactionNotFoundException struct {
	Message string
}

func (e *TransactionNotFoundException) Error() string {
	return fmt.Sprintf("Not able to find the transaction specified: %s", (*e).Message)
}

type TransactionCreationException struct {
	Message string
}

func (e *TransactionCreationException) Error() string {
	return fmt.Sprintf("Not able to create the transaction: %s", (*e).Message)
}

type InsufficientBalanceException struct {
	Message string
}

func (e *InsufficientBalanceException) Error() string {
	return fmt.Sprintf("Insufficient Balance: %s", (*e).Message)
}

type AccountCreationException struct {
	Message string
}

func (e *AccountCreationException) Error() string {
	return fmt.Sprintf("Not able to create account: %s", (*e).Message)
}

type AccountNotFoundException struct {
	Message string
}

func (e *AccountNotFoundException) Error() string {
	return fmt.Sprintf("Not able to find the account: %s", (*e).Message)
}

type AccountUpdateException struct {
	Message string
}

func (e *AccountUpdateException) Error() string {
	return fmt.Sprintf("Not able to update the account: %s", (*e).Message)
}

type DataBaseConnectionException struct {
	Message string
}

func (e *DataBaseConnectionException) Error() string {
	return fmt.Sprintf("Failed to do database connection: %s", (*e).Message)
}
