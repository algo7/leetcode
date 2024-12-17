package ispalindrome

import (
	"slices"
	"strconv"
	"strings"
)

func IsPalindromeString(x int) bool {

	ns := strconv.Itoa(x)

	nsao := strings.Split(ns, "")
	nsa := strings.Split(ns, "")

	slices.Reverse(nsao)

	for i := range nsao {

		if nsa[i] != nsao[i] {
			return false
		}
	}

	return true
}

func IsPalindromeStringShort(x int) bool {

	ns := strconv.Itoa(x)
	nsl := len(ns)

	for i := 0; i < nsl/2; i++ {
		if ns[i] != ns[nsl-1-i] {
			return false
		}
	}

	return true
}

func IsPalindromeMod(x int) bool {

	var reversed int
	tmp := x

	for tmp > 0 {

		reversed = reversed*10 + tmp%10 // tmp%10 gets us the last digit

		tmp = tmp / 10 // this get us the remaining digits
	}

	return x == reversed

}
