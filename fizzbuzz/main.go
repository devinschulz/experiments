package main

import (
	"fmt"

	"github.com/devinschulz/experiments/fizzbuzz/fizzbuzz"
)

func main() {
	for _, item := range fizzbuzz.Run(30) {
		fmt.Println(item)
	}
}
