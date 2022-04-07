package main

import (
	"bufio"
	"fmt"
	"os"
	cp "persofin/src/command_parser"
	cc "persofin/src/commons"
	logger "persofin/src/logger"
)

func main() {
	var input string = ""
	reader := bufio.NewReader(os.Stdin) // creating reader

	logger.SetupLogging()
	logger.PrintLog(logger.INFO, "Initiating persofin..")

	for {
		fmt.Print(cc.PROMT)
		input, _ = reader.ReadString('\n') // getting user input
		if cp.CommandProcessor(input) == cc.FAILURE {
			logger.PrintLog(logger.INFO, "Persofin termination initiated...")
			break
		}
	}
	logger.PrintLog(logger.INFO, "Persofin terminated")
	logger.LoggingCleanup()
}
