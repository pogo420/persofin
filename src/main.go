package main

import (
	"fmt"
	cp "persofin/src/command_parser"
	cc "persofin/src/commons"
)

func main() {
	var input string = ""
	for {
		fmt.Print(cc.PROMT)
		fmt.Scanln(&input)
		if cp.CommandProcessor(input) == -1 {
			break
		}
	}
}
