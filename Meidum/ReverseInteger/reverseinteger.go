package reverseinteger

import (
	"math"
)

func Reverse(x int) int {

	tmp, reversed := x, 0

	// check if it's negative
	isNegative := false
	if math.Signbit(float64(x)) {
		tmp = -x
		isNegative = true
	}

	for tmp > 0 {

		reversed = reversed*10 + tmp%10
		tmp = tmp / 10

	}

	if isNegative {

		if -reversed <= math.MinInt32 {
			return 0
		} else {
			return -reversed
		}

	}

	if reversed >= math.MaxInt32 {
		return 0
	}

	return reversed

}
