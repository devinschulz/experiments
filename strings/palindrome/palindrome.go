package strings

import (
	"regexp"
	str "strings"
)

// Determine if a string is a valid palindrome.
func Palindrome(s string) bool {
	sanitizeString(&s)
	r := str.Split(str.ToLower(str.Trim(s, "")), "")

	for i, x := 0, len(r)-1; i < len(r)/2; i, x = i+1, x-1 {
		if r[i] != r[x] {
			return false
		}
	}
	return true
}

func sanitizeString(s *string) {
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	*s = reg.ReplaceAllString(*s, "")
}
