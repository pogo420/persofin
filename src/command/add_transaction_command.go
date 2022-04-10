package command

// TODO: Add account exists logic
// TODO: Add date validation logic(In command package scope)
// TODO: Handle wrong issues
// TODO add test case
import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	"strconv"
	"strings"
)

type AddTransactionCommand struct {
}

// function to split accounts and transaction value
func (atc *AddTransactionCommand) splitAccounts(flags string) (cmm.Date, cmm.AccountName, cmm.TansactionValue) {
	splitString := strings.Fields(flags)
	date_ := cmm.Date(splitString[0])
	accountName := cmm.AccountName(splitString[1])
	tranValue, err := strconv.Atoi(splitString[2])
	if err != nil {
		return "", "", 0
	} else if len(splitString) <= 2 {
		return "", "", 0
	} else {
		return date_, accountName, cmm.TansactionValue(tranValue)
	}
}
func (atc *AddTransactionCommand) Description() string {
	return fmt.Sprintf("Command to add transaction into an account. Example: %s <date> <acc_name> <transaction_value>", cmm.ADD_TRANSACTION_COMMAND)
}

func (atc *AddTransactionCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {
	_, _, _ = (*atc).splitAccounts(flags)
	return cmm.FAILURE
}
