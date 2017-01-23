package algorithms_test

import (
	"github.com/devinschulz/experiments/algorithms/sort"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, pair := range tests {
		v := algorithms.InsertionSort(pair.value)
		if !reflect.DeepEqual(v, pair.result) {
			t.Error(
				"for", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func BenchmarkInsertionSortHundred(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.InsertionSort(hundred)
	}
}

func BenchmarkInsertionSortThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.InsertionSort(thousand)
	}
}

func BenchmarkInsertionSortTenThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.InsertionSort(tenThousand)
	}
}
