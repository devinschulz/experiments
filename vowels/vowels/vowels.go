package vowels

import "strings"

var vowels = []string{"a", "e", "i", "o", "u"}

func Count(str string) int {
	str = strings.ToLower(str)
	sum := 0
	for _, vowel := range vowels {
		sum += strings.Count(str, vowel)
	}
	return sum
}
