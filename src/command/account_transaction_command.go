package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
)

type AccountTransactionCommand struct {
}

func (atc *AccountTransactionCommand) Description() string {
	return fmt.Sprintf("Command to move value from one account to another. Example: %s <acc_name_1> <acc_name_2> <transaction_value>", cmm.ACCOUNT_TRANSACTION_COMMAND)

}

func (atc *AccountTransactionCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {
	return cmm.FAILURE
}
