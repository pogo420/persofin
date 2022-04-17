package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	logger "persofin/src/logger"
)

type GetAllBalanceCommand struct {
}

func (gabc *GetAllBalanceCommand) Description() string {
	return fmt.Sprintf("Command to get balances of all account. Example: %s", cmm.GET_ALL_BALANCE_COMMAND)
}

func (gabc *GetAllBalanceCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {
	logger.PrintLog(logger.INFO, fmt.Sprintf("Executing command: %s", cmm.GET_ALL_BALANCE_COMMAND))

	balance := dbInterface.GetAllAccountsBalance()
	logger.PrintLog(logger.INFO, fmt.Sprintf("Account balances %s", balance))
	fmt.Printf("%sAccount balances %s\n", cmm.PROMT, balance)
	return cmm.SUCCESS
}
