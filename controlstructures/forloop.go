package main

import (
	"fmt"
	a "github.com/benhalstead/gotraining/annotation"
)

// There are no while or do while loops in Go. The for statement is flexible enough to simulate those
func main() {

	// The standard 'for' statement follows the C convention with an initialisation, a test and an increment/decrement
	a.Section("For statement")

	for i := 0; i < 2; i++ {
		fmt.Printf("Counter is %d\n", i)
	}

	// If you just have a test, you can simulate a loop
	a.Section("For statement as a loop")
	n := 3

	for n >= 0 {
		fmt.Printf("Counter is %d\n", n)
		n--
	}

	// continue 'skips' an iteration in a loop
	a.Section("Using continue in a for loop")

	for i := 0; i < 5; i++ {

		if i == 2 || i == 4 {
			fmt.Printf("Ignoring %d\n", i)
			continue
		}

		fmt.Printf("Counter is %d\n", i)
	}

	// break exits the current structure early
	a.Section("Breaking out of a for loop")

	for i := 0; i < 5; i++ {

		if i == 2 {
			fmt.Printf("Breaking at %d\n", i)
			break
		}

		fmt.Printf("Counter is %d\n", i)
	}

	// Labels can be used to break to specific a.Sections of code - useful for nested loops
	a.Section("Breaking to a label")

MyLabel:
	for i := 0; i < 2; i++ {

		fmt.Printf("Outer is %d\n", i)

		for j := 0; j < 3; j++ {

			if i == 1 && j == 2 {
				fmt.Printf("Breaking to label at %d:%d\n", i, j)
				break MyLabel
			}

			fmt.Printf("Inner is %d\n", j)

		}

	}
}
