package romantointeger

import (
	"strings"
)

func RomanToIntFirst(s string) int {

	intMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	splitted := strings.Split(s, "")
	strLen := len(splitted)
	num := 0

	for i := 0; i < strLen; i++ {
		// handle no. 4
		if splitted[i] == "I" && i+1 < strLen && splitted[i+1] == "V" {
			num = num + 4
			i++
			continue
		}

		// handle no. 9
		if splitted[i] == "I" && i+1 < strLen && splitted[i+1] == "X" {
			num = num + 9
			i++
			continue
		}

		// handle 40
		if splitted[i] == "X" && i+1 < strLen && splitted[i+1] == "L" {
			num = num + 40
			i++
			continue
		}

		// handle 90
		if splitted[i] == "X" && i+1 < strLen && splitted[i+1] == "C" {
			num = num + 90
			i++
			continue
		}

		// handle 400
		if splitted[i] == "C" && i+1 < strLen && splitted[i+1] == "D" {
			num = num + 400
			i++
			continue
		}

		// handle 900
		if splitted[i] == "C" && i+1 < strLen && splitted[i+1] == "M" {
			num = num + 900
			i++
			continue
		}

		v, ok := intMap[splitted[i]]
		if ok {
			num = num + v
		}
	}

	return num
}
