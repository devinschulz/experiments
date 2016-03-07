package fizzbuzz

import "testing"

var test = []string{
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
	"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz",
}

func TestRun(t *testing.T) {
	v := Run(75)
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
