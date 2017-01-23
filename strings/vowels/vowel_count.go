package strings

import str "strings"

var vowels = [...]string{"a", "e", "i", "o", "u"}

func VowelCount(string string) int {
	string = str.ToLower(string)
	sum := 0
	for _, vowel := range vowels {
		sum += str.Count(string, vowel)
	}
	return sum
}
