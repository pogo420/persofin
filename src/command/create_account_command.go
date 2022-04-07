package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	logger "persofin/src/logger"
)

type CreateAccountCommand struct {
}

func (cac *CreateAccountCommand) Description() string {
	return fmt.Sprintf("Command to create account. Example: %s <acc_name>", cmm.CREATE_ACCOUNT_COMMAND)
}

func (cac *CreateAccountCommand) Execute(flags string) int {
	logger.PrintLog(logger.INFO, "Getting db interface")
	dbInterface := dbi.GetDbInterface() // getting interface

	// Handling account exists case
	if dbInterface.AccountExists(cmm.AccountName(flags)) {
		logger.PrintLog(logger.ERROR, "Error in creating account, Account Exists")
		return cmm.FAILURE
	}

	// creating account logic
	logger.PrintLog(logger.INFO, fmt.Sprintf("Getting db interface: %s", dbInterface.GetInterfaceVersion()))
	response := dbInterface.CreateAccount(cmm.AccountName(flags))

	// Handling response
	if response == cmm.FAILURE {
		logger.PrintLog(logger.ERROR, "Error in creating account")
	} else {
		logger.PrintLog(logger.INFO, "Creating account successful")
	}
	return response
}
