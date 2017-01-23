package strings_test

import (
	"testing"
	"github.com/devinschulz/experiments/strings/vowels"
)

type test struct {
	value string
	count int
}

var tests = []test{
	{"The quick brown fox jumps over the lazy dog.", 11},
	{"How quickly daft jumping zebras vex.", 9},
	{"Sphinx of black quartz: judge my vow.", 8},
	{"Mr. Jock, TV quiz PhD, bags FEW lynx.", 5},
}

func TestCount(t *testing.T) {
	for _, pair := range tests {
		v := strings.VowelCount(pair.value)
		if v != pair.count {
			t.Error(
				"for", pair.value,
				"expected", pair.count,
				"got", v,
			)
		}
	}
}
