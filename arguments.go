package utilities

import (
	"os"

	errortools "github.com/leapforce-libraries/go_errortools"
)

func GetArguments(arguments ...*string) *errortools.Error {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < len(arguments) {
		return errortools.ErrorMessage("Too little arguments passed.")
	}

	for index, argument := range arguments {
		(*argument) = argsWithoutProg[index]
	}

	return nil
}
