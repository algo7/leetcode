package fizzbuzz

import (
	"math"
	"strconv"
)

func FizzBuzzFirst(n int) []string {

	nums := make([]float64, n)
	numString := []string{}

	for i := range nums {
		nums[i] = float64(i) + float64(1)
	}

	for _, n := range nums {
		if math.Mod(n, 3) == 0 && math.Mod(n, 5) == 0 {
			numString = append(numString, "FizzBuzz")
		} else if math.Mod(n, 3) == 0 {
			numString = append(numString, "Fizz")

		} else if math.Mod(n, 5) == 0 {
			numString = append(numString, "Buzz")
		} else {
			numString = append(numString, strconv.FormatFloat(n, 'f', 0, 64))
		}
	}

	return numString
}

func FizzBuzzSecond(n int) []string {

	out := []string{}

	for i := 1; i <= n; i++ {

		switch {

		case i%15 == 0:
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, strconv.Itoa(i))

		}
	}

	return out

}
