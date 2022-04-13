package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	logger "persofin/src/logger"
	"strconv"
	"strings"
)

type AccountTransactionCommand struct {
}

// function to split accounts and transaction value
func (atc *AccountTransactionCommand) splitFlags(flags string) (cmm.Date, cmm.AccountName, cmm.AccountName, cmm.TansactionValue) {

	splitString := strings.Fields(flags)

	if len(splitString) <= 3 { // invalid flags
		return "", "", "", 0
	}

	// segregating data
	date_ := cmm.Date(splitString[0])
	accountNameFrom := cmm.AccountName(splitString[1])
	accountNameTo := cmm.AccountName(splitString[2])
	tranValue, err := strconv.Atoi(splitString[3])

	if err != nil {
		return "", "", "", 0
	} else {
		return date_, accountNameFrom, accountNameTo, cmm.TansactionValue(tranValue)
	}
}

func (atc *AccountTransactionCommand) Description() string {
	return fmt.Sprintf("Command to move value from one account to another. Example: %s <date> <acc_name_1> <acc_name_2> <transaction_value>", cmm.ACCOUNT_TRANSACTION_COMMAND)

}

func (atc *AccountTransactionCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {

	logger.PrintLog(logger.INFO, fmt.Sprintf("Executing command: %s", cmm.ADD_TRANSACTION_COMMAND))

	date_, accountNameFrom, accountNameTo, transactionValue := (*atc).splitFlags(flags)

	if date_ == "" || !isValidateDate(date_) { // missing date
		logger.PrintLog(logger.ERROR, "Issue in validating date")
		return cmm.FAILURE

	} else if accountNameFrom == "" || accountNameTo == "" { // missing account name
		logger.PrintLog(logger.ERROR, "Issue in Account Name")
		return cmm.FAILURE

	} else if !dbInterface.AccountExists(cmm.AccountName(accountNameFrom)) || !dbInterface.AccountExists(cmm.AccountName(accountNameTo)) { // if accounts does not exist
		logger.PrintLog(logger.ERROR, "Account name does not exist")
		return cmm.FAILURE

	} else {
		// adding transaction
		logger.PrintLog(logger.INFO, fmt.Sprintf("Adding Transaction from %s to to %s", accountNameFrom, accountNameTo))
		response := dbInterface.AccountTransaction(date_, accountNameFrom, accountNameTo, transactionValue)

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
