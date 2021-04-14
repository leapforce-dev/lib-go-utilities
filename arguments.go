package utilities

import (
	"fmt"
	"os"

	errortools "github.com/leapforce-libraries/go_errortools"
)

func GetArguments(required *int, arguments ...*string) *errortools.Error {
	argsWithoutProg := os.Args[1:]

	_required := len(arguments)
	if required != nil {
		if *required > len(arguments) {
			return errortools.ErrorMessage(fmt.Sprintf("%v arguemnts passed but required %v", len(arguments), *required))
		}
		_required = *required
	}

	if len(argsWithoutProg) < _required {
		return errortools.ErrorMessage("Too little arguments passed.")
	}

	for index, arg := range argsWithoutProg {
		if index >= _required {
			break
		}
		*(arguments[index]) = arg
	}

	return nil
}
