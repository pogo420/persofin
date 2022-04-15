package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
)

type GetBalanceCommand struct {
}

func (gbc *GetBalanceCommand) Description() string {
	return fmt.Sprintf("Command to get current balance of an account. Example: %s <acc_name>", cmm.GET_BALANCE_COMMAND)

}

func (gbc *GetBalanceCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {
	return cmm.FAILURE
}
