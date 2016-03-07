package main

import (
	"fmt"

	"github.com/devinschulz/experiments/algorithms/algorithms"
)

func main() {
	fmt.Println("Bubble Sort")
	fmt.Println(algorithms.BubbleSort([]int{1, 3, 5, 29, 27, 2, 3, 19, 5}))
	fmt.Println(algorithms.BubbleSort([]int{19, 28, 10, 20, 18, 10, 10, 200, 89}))
	fmt.Println(algorithms.BubbleSort([]int{90, 129, 100, 29, 910, 237, 192, 191, 271, 38}))

	fmt.Println("Insertion Sort")
	fmt.Println(algorithms.InsertionSort([]int{1, 3, 5, 29, 27, 2, 3, 19, 5}))
	fmt.Println(algorithms.InsertionSort([]int{19, 28, 10, 20, 18, 10, 10, 200, 89}))
	fmt.Println(algorithms.InsertionSort([]int{90, 129, 100, 29, 910, 237, 192, 191, 271, 38}))
}
