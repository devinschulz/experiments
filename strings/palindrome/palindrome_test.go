package strings_test

import (
	"testing"
	"github.com/devinschulz/experiments/strings/palindrome"
)

type test struct {
	value  string
	result bool
}

var tests = []test{
	{"racecar", true},
	{"A car, a man, a maraca.", true},
	{"The quick brown fox jumps over the lazy dog.", false},
	{"Eva, can I see bees in a cave?", true},
	{"Top spots", false},
	{"Rotator", true},
}

func TestPalindrome(t *testing.T) {
	for _, pair := range tests {
		v := strings.Palindrome(pair.value)
		if v != pair.result {
			t.Error(
				"for", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
