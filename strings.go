package utilities

import (
	"bytes"
	"regexp"
	"strings"
)

// StringSliceContains checks whether a string is present in a slice
func StringSliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// IsLetter checks whether string contains of just letters
func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func NormalizeString(s string, removeSymbols bool, removeRegex *string) string {
	s = strings.Trim(s, " ")

	reader := bytes.NewReader([]byte(s))

	var result string

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		switch r {
		case 138:
			result += "S"
		case 140:
			result += "OE"
		case 142:
			result += "Z"
		case 154:
			result += "s"
		case 156:
			result += "oe"
		case 158:
			result += "z"
		case 159:
			result += "Y"
		case 192, 193, 194, 195, 196, 197:
			result += "A"
		case 198:
			result += "AE"
		case 199:
			result += "C"
		case 200, 201, 202, 203:
			result += "E"
		case 204, 205, 206, 207:
			result += "I"
		case 208:
			result += "D"
		case 209:
			result += "N"
		case 210, 211, 212, 213, 214, 216:
			result += "O"
		case 217, 218, 219, 220:
			result += "U"
		case 221:
			result += "Y"
		case 222:
			result += "p"
		case 223:
			result += "ss"
		case 224, 225, 226, 227, 228, 229:
			result += "a"
		case 230:
			result += "ae"
		case 231:
			result += "c"
		case 232, 233, 234, 235:
			result += "e"
		case 236, 237, 238, 239:
			result += "i"
		case 240:
			result += "d"
		case 241:
			result += "n"
		case 242, 243, 244, 245, 246, 248:
			result += "o"
		case 249, 250, 251, 252:
			result += "u"
		case 253:
			result += "y"
		default:
			result += string(r)

		}
	}

	if removeRegex != nil {
		re := regexp.MustCompile(*removeRegex)
		result = re.ReplaceAllString(result, "")
	}

	if removeSymbols {
		re := regexp.MustCompile(`[^\w|\s]`)
		result = re.ReplaceAllString(result, "")
	}

	return result
}
