package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	logger "persofin/src/logger"
	"strconv"
	"strings"
)

type AddTransactionCommand struct {
}

// function to split accounts and transaction value
func (atc *AddTransactionCommand) splitAccounts(flags string) (cmm.Date, cmm.AccountName, cmm.TansactionValue) {

	splitString := strings.Fields(flags)

	if len(splitString) <= 2 { // invalid flags
		return "", "", 0
	}

	// segregating data
	date_ := cmm.Date(splitString[0])
	accountName := cmm.AccountName(splitString[1])
	tranValue, err := strconv.Atoi(splitString[2])

	if err != nil {
		return "", "", 0
	} else {
		return date_, accountName, cmm.TansactionValue(tranValue)
	}
}
func (atc *AddTransactionCommand) Description() string {
	return fmt.Sprintf("Command to add transaction into an account. Example: %s <date> <acc_name> <transaction_value>", cmm.ADD_TRANSACTION_COMMAND)
}

func (atc *AddTransactionCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {

	logger.PrintLog(logger.INFO, fmt.Sprintf("Executing command: %s", cmm.ADD_TRANSACTION_COMMAND))

	date_, accountName, transactionValue := (*atc).splitAccounts(flags)

	if date_ == "" || !isValidateDate(date_) { // missing date
		logger.PrintLog(logger.ERROR, "Issue in validating date")
		return cmm.FAILURE

	} else if accountName == "" { // missing account name
		logger.PrintLog(logger.ERROR, "Issue in Account Name")
		return cmm.FAILURE

	} else if !dbInterface.AccountExists(cmm.AccountName(flags)) { // if accounts does not exist
		logger.PrintLog(logger.ERROR, "Account name does not exist")
		return cmm.FAILURE

	} else {
		// adding transaction
		logger.PrintLog(logger.INFO, fmt.Sprintf("Adding Transaction into %s", accountName))
		response := dbInterface.AddTransaction(date_, accountName, transactionValue)

		// handling response
		if response == cmm.FAILURE {
			logger.PrintLog(logger.ERROR, "Issue in addng transaction")
			return cmm.FAILURE

		} else {
			logger.PrintLog(logger.INFO, "Adding transaction successful")
			return cmm.SUCCESS

		}
	}
}
