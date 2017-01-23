package strings_test

import (
	"testing"
	"github.com/devinschulz/experiments/strings/fizzbuzz"
)

var test = []string{
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
}

func TestRun(t *testing.T) {
	v := strings.FizzBuzz(75)
	for i, d := range test {
		if d != v[i] {
			t.Error(
				"expected", d,
				"got", v[i],
				"at", i,
			)
		}
	}
}
