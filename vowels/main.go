package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/devinschulz/experiments/vowels/vowels"
)

func main() {
	fmt.Print("Enter a sentence: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}

	fmt.Printf("\"%v\" has %d vowels!\n", scanner.Text(), vowels.Count(scanner.Text()))
}
