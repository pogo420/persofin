package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	logger "persofin/src/logger"
)

type GetBalanceCommand struct {
}

func (gbc *GetBalanceCommand) Description() string {
	return fmt.Sprintf("Command to get current balance of an account. Example: %s <acc_name>", cmm.GET_BALANCE_COMMAND)

}

func (gbc *GetBalanceCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {
	logger.PrintLog(logger.INFO, fmt.Sprintf("Executing command: %s", cmm.GET_BALANCE_COMMAND))
	logger.PrintLog(logger.INFO, "Checking for account exists")
	// Handling account exists case
	if !dbInterface.AccountExists(cmm.AccountName(flags)) {
		logger.PrintLog(logger.ERROR, "Account does not Exists")
		return cmm.FAILURE
	}

	balance := dbInterface.GetAccountBalance(cmm.AccountName(flags))
	logger.PrintLog(logger.INFO, fmt.Sprintf("Account balance of %s is %d", flags, balance))
	fmt.Printf("%sAccount balance of %s is %d\n", cmm.PROMT, flags, balance)
	return cmm.SUCCESS
}
