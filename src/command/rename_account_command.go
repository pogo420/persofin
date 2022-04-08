package command

import (
	"fmt"
	cmm "persofin/src/commons"
)

type RenameAccountCommand struct {
}

func (rac *RenameAccountCommand) Description() string {
	return fmt.Sprintf("Command to rename account. Example: %s <acc_name> <account_name1> <account_name2>", cmm.RENAME_ACCOUNT_COMMAND)
}

func (rac *RenameAccountCommand) Execute(flags string) int {
	fmt.Println(flags)
	return -1
}
