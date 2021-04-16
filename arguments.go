package utilities

import (
	"fmt"
	"os"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
)

func GetArguments(required *int, arguments ...*string) (*map[string]string, *errortools.Error) {
	argsWithoutProg := os.Args[1:]
	var args []string = []string{}
	var prefixedArgs map[string]string = make(map[string]string)

	// first extract arguments passed with - prefix
	for _, arg := range argsWithoutProg {
		if strings.HasPrefix(arg, "-") {
			if len(arg) == 1 {
				return nil, errortools.ErrorMessage("Invalid argument '-'")
			}
			if !IsLetter(arg[1:2]) {
				return nil, errortools.ErrorMessage(fmt.Sprintf("Invalid prefix '%s'", arg[:2]))
			}
			prefixedArgs[arg[1:2]] = arg[2:]
		} else {
			args = append(args, arg)
		}
	}

	_required := len(arguments)
	if required != nil {
		if *required > len(arguments) {
			return nil, errortools.ErrorMessage(fmt.Sprintf("%v arguments passed but required %v", len(arguments), *required))
		}
		_required = *required
	}

	if len(args) < _required {
		return nil, errortools.ErrorMessage("Too little arguments passed.")
	}

	for index, arg := range args {
		if index >= _required {
			break
		}
		*(arguments[index]) = arg
	}

	return &prefixedArgs, nil
}
