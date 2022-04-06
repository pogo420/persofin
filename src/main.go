package main

import (
	"fmt"
	cp "persofin/src/command_parser"
	cc "persofin/src/commons"
	logger "persofin/src/logger"
)

func main() {
	var input string = ""
	logger.SetupLogging()
	logger.PrintLog(logger.INFO, "Initiating persofin..")
	for {
		fmt.Print(cc.PROMT)
		fmt.Scanln(&input)
		if cp.CommandProcessor(input) == cc.FAILURE {
			logger.PrintLog(logger.INFO, "Persofin termination initiated...")
			break
		}
	}
	logger.PrintLog(logger.INFO, "Persofin terminated")
	logger.LoggingCleanup()
}
