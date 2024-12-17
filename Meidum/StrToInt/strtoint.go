package strtoint

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func MyAtoi(s string) int {

	// slice of valid numbers (you can use map ofc)
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// check for empty string, or any other length-1 string with non-numeric character
	if s == "" || len(s) == 1 && !slices.Contains(digits, s) {
		return 0
	}

	// trim the space and split the string into a string slice
	spaceTrimmed := strings.TrimSpace(s)
	sa := strings.Split(spaceTrimmed, "")

	// check for slice with the length of 0. this will happen if the input string is only space(s)
	if len(sa) == 0 {
		return 0
	}

	// check for when it's indicated that the number is strictly positive: +12
	// and not followed by other non-numerical characters: +a12
	isExplicitPositive := false
	if sa[0] == "+" && slices.Contains(digits, sa[1]) {
		isExplicitPositive = true
		sa = sa[1:]
	}

	// check if the number is negative
	isNegative := false
	if !isExplicitPositive {
		if sa[0] == "-" {
			// take a note that it's negative and remove the - sign and remove it to avoid interfering with the logic after
			isNegative = true
			sa = sa[1:]
		}
	}

	// check for when after removing the +/- the first character is still non-numeric
	if !slices.Contains(digits, sa[0]) {
		return 0
	}

	// var to hold the index
	stop := 0

	for i, str := range sa {

		// stop after encontering the 1st non-numberic character and take note of its index
		isNum := slices.Contains(digits, str)
		if !isNum {
			stop = i
			break
		}
	}

	// if non-numeric character is encountered the take only the part of the slice that comes before it
	if stop > 0 {
		sa = sa[:stop]
	}

	// merge the string slice back to a string
	finalStr := strings.Join(sa, "")

	// convert it to number.
	// there must be a way not to call the actual atoi
	x, _ := strconv.Atoi(finalStr)

	// check if the resulting int is out of bound
	if math.Abs(float64(x)) > math.MaxInt32 {

		switch isNegative {
		case true:
			return math.MinInt32
		case false:
			return math.MaxInt32
		}

	} else {
		switch isNegative {
		case true:
			return -x
		case false:
			return x
		}
	}

	return 0
}
