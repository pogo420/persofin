package command

import (
	"fmt"
	cmm "persofin/src/commons"
	dbi "persofin/src/db_interface"
	logger "persofin/src/logger"
	"strings"
)

type RenameAccountCommand struct {
}

// function to split accounts
func (rac *RenameAccountCommand) splitAccounts(flags string) (string, string) {
	splitString := strings.Fields(flags)
	if len(splitString) <= 1 {
		return "", ""
	} else {
		return splitString[0], splitString[1]
	}
}

func (rac *RenameAccountCommand) Description() string {
	return fmt.Sprintf("Command to rename account. Example: %s <acc_name> <account_name_old> <account_name_new>", cmm.RENAME_ACCOUNT_COMMAND)
}

func (rac *RenameAccountCommand) Execute(dbInterface dbi.BaseDbInterface, flags string) int {
	logger.PrintLog(logger.INFO, fmt.Sprintf("Executing command: %s", cmm.RENAME_ACCOUNT_COMMAND))
	logger.PrintLog(logger.INFO, "Extracting account names")

	// extracting accounts information
	account_old, account_new := (*rac).splitAccounts(flags)
	logger.PrintLog(logger.INFO, fmt.Sprintf("Accounts extracted, old: %s, new: %s", account_old, account_new))

	// checking for nonexistance of old account exists and existance of new acount
	if !dbInterface.AccountExists(cmm.AccountName(account_old)) || dbInterface.AccountExists(cmm.AccountName(account_new)) {
		logger.PrintLog(logger.ERROR, "Error in renaming account, new Account name might Exists or old account name might not exist")
		return cmm.FAILURE
	}
	// renaming account via api call
	response := dbInterface.RenameAccount(cmm.AccountName(account_old), cmm.AccountName(account_new))

	// response handling
	if response == cmm.FAILURE {
		logger.PrintLog(logger.ERROR, "Issue in renaming account")
		return cmm.FAILURE
	}
	logger.PrintLog(logger.INFO, "Issue in renaming account")
	return cmm.SUCCESS
}
