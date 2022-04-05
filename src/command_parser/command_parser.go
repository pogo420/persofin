package command_parser

import (
	"fmt"
	cc "persofin/src/commons"
)

// Function to check valid commands
func findInCommandList(command string) bool {
	for key := range cc.Commands {
		if key == command {
			return true
		}
	}
	return false
}

// Function to print list of commands
func listCommand() {
	fmt.Println("Listing commands..")
	for key, value := range cc.Commands {
		fmt.Printf("command:%s - description:%s\n", key, value.Description())
	}
}

// Public function to process command
// for invalid command - retun list of command and function returns 0
// for exit command bye function returns -1
// For valid command function retuns 0 or -1 based on success/failure
func CommandProcessor(input string) int {
	if !findInCommandList(input) {
		listCommand()
	} else if cc.Commands[input].Execute() == -1 { // command is executed and return code is checked
		return -1
	}
	return 0
}
