package utils

import "os"

const Test_mode_flag string = "PERSOFIN_TEST"

func ReadEnvVariable(variableName string) string {
	// function to return env variable values
	return os.Getenv(variableName)
}

func IsTestMode() bool {
	// function to check the test mode
	if ReadEnvVariable(Test_mode_flag) == "1" {
		return true
	} else {
		return false
	}
}
