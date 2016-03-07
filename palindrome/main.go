package main

import (
	"fmt"

	"github.com/devinschulz/experiments/palindrome/palindrome"
)

func main() {
	fmt.Println(palindrome.Valid("racecar"))                                      // true
	fmt.Println(palindrome.Valid("A car, a man, a maraca."))                      // true
	fmt.Println(palindrome.Valid("The quick brown fox jumps over the lazy dog.")) // false
	fmt.Println(palindrome.Valid("A Santa dog lived as a devil God at NASA."))    // true
}
