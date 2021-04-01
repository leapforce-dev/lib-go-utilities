package utilities

import (
	"fmt"
	"strings"
)

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func SplitAddress(address string) (street string, house string) {
	address = strings.Trim(address, " ")
	if len(address) == 0 {
		return "", ""
	}

	length := len(address)
	fields := strings.Fields(address)
	size := len(fields)

	if size <= 1 {
		return address, ""
	}
	last := fields[size-1]
	penult := fields[size-2]
	if isDigit(last[0]) {
		isdig := isDigit(penult[0])
		if size > 2 && isdig && !strings.HasPrefix(penult, "194") {
			house = fmt.Sprintf("%s %s", penult, last)
		} else {
			house = last
		}
	} else if size > 2 {
		house = fmt.Sprintf("%s %s", penult, last)
	}
	street = strings.TrimRight(address[:length-len(house)], " ")
	return
}
