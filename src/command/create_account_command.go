package command

import (
	"fmt"
	cmm "persofin/src/commons"
)

type CreateAccountCommand struct {
}

func (cac *CreateAccountCommand) Description() string {
	return fmt.Sprintf("Command to create account. Example: %s <acc_name>", cmm.CREATE_ACCOUNT_COMMAND)
}

func (cac *CreateAccountCommand) Execute(flags string) int {
	fmt.Printf("flags %s\n", flags)
	return -1
}
