package strings

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func FizzBuzz(limit int) []string {
	v := []string{}
	for i := 1; i <= limit; i++ {
		threes := i%3 == 0
		fives := i%5 == 0
		switch {
		case threes && fives:
			v = append(v, fizz+buzz)
		case threes:
			v = append(v, fizz)
		case fives:
			v = append(v, buzz)
		}
	}
	return v
}
