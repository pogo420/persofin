package command_mapper

import (
	cm "persofin/src/command"
	co "persofin/src/commons"
)

// All commands to struct mapping
var Commands = map[string]cm.BaseCommand{
	co.EXIT_COMMAND:                &cm.ExitCommand{},
	co.CREATE_ACCOUNT_COMMAND:      &cm.CreateAccountCommand{},
	co.RENAME_ACCOUNT_COMMAND:      &cm.RenameAccountCommand{},
	co.ADD_TRANSACTION_COMMAND:     &cm.AddTransactionCommand{},
	co.ACCOUNT_TRANSACTION_COMMAND: &cm.AccountTransactionCommand{},
}
