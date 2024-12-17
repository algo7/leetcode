package main

import (
	"log"

	reverseinteger "github.com/algo7/leetcode/Meidum/ReverseInteger"
)

func main() {
	r := reverseinteger.Reverse(-2147483648)
	log.Println(r)
}
