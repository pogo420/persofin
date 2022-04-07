package command

import (
	"fmt"
	cmm "persofin/src/commons"
)

// Exit command logic
type ExitCommand struct {
}

func (ex *ExitCommand) Description() string {
	return fmt.Sprintf("Command to exit program. Example: %s", cmm.EXIT_COMMAND)
}

func (ex *ExitCommand) Execute(input string) int {
	return -1
}
