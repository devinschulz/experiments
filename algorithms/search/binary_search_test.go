package algorithms_test

import (
	"github.com/devinschulz/experiments/algorithms/search"
	"testing"
)

var tests = []struct {
	searchKey int
	result    int
}{
	{90, 4},
	{100, 5},
	{250, 8},
	{180, 7},
	{0, -1},
	{101, -1},
	{10000, -1},
}

func TestBinarySearch(t *testing.T) {
	searchList := []int{10, 15, 25, 45, 90, 100, 140, 180, 250}
	for _, pair := range tests {
		searchIndex := algorithms.BinarySearch(searchList, pair.searchKey)
		if searchIndex != pair.result {
			t.Error(
				"for", pair.searchKey,
				"expected", pair.result,
				"got", searchIndex,
			)
		}
	}
}
