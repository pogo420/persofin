package command_parser

import (
	"fmt"
	cc "persofin/src/commons/command_mapper"
	"strings"
)

// Function to split command and flag
func splitCommandFlag(input string) (string, string) {
	splitInput := strings.Fields(input)

	if len(splitInput) == 0 {
		return "", ""
	} else if len(splitInput) == 1 {
		return splitInput[0], ""
	} else {
		return splitInput[0], strings.Join(splitInput[1:], " ")
	}
}

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
	command, flags := splitCommandFlag(input)
	if !findInCommandList(command) {
		listCommand()
	} else if cc.Commands[command].Execute(flags) == -1 { // command is executed and return code is checked
		return -1
	}
	return 0
}
