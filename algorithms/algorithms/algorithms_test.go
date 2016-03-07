package algorithms

import (
	"reflect"
	"testing"
)

type test struct {
	value  []int
	result []int
}

var tests = []test{
	{[]int{0, 8, 9, 2, 1, 6, 5, 10, 3, 4, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{82, 29, 21, 61, 18, 4, 8, 10, 1, 6, 81, 95}, []int{1, 4, 6, 8, 10, 18, 21, 29, 61, 81, 82, 95}},
	{[]int{100, 2810, 28, 1903, 839, 282}, []int{28, 100, 282, 839, 1903, 2810}},
	{[]int{-8, -19, -20, 2, 0, 5, 12, 22}, []int{-20, -19, -8, 0, 2, 5, 12, 22}},
}

func TestBubbleSort(t *testing.T) {
	for _, pair := range tests {
		arr := BubbleSort(pair.value)
		if !reflect.DeepEqual(arr, pair.result) {
			t.Error(
				"for", pair.value,
				"expected", pair.result,
				"got", arr,
			)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, pair := range tests {
		v := IntersectionSort(pair.value)
		if !reflect.DeepEqual(v, pair.result) {
			t.Error(
				"for", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
