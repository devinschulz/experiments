package fizzbuzz

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func Run(limit int) []string {
	v := make([]string, 0)
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
