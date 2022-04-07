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

func (cac *CreateAccountCommand) Execute(input string) int {

	fmt.Printf("")
	return -1
}
