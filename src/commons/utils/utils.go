package utils

import "os"

func ReadEnvVariable(variableName string) string {
	// function to return env variable values
	return os.Getenv(variableName)
}
