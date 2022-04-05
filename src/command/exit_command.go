package command

// Exit command logic
type ExitCommand struct {
}

func (ex *ExitCommand) Description() string {
	return "Command to exit program"
}

func (ex *ExitCommand) Execute() int {
	return -1
}
