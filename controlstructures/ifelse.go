package main

import (
	"fmt"
	a "github.com/benhalstead/gotraining/tutorial"
)

// if/else if/else block is common to most C-syntax languages
func main() {

	a.Section("Basic if statement")

	n := 2

	if n == 2 {
		fmt.Println("Match")
	}

	// But the layout of the clause keywords is enforced by the compiler - else and else if must be on the same
	// line as the closing brace for the previous clause
	a.Section("if/else if/else statement")

	if n < 0 {
		fmt.Println("Negative")
	} else if n > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}
}
