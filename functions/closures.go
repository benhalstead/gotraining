package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	// In Go the syntax for closures is shared with lambdas/anonymous functions

	tutorial.Section("Anonymous function/lambda expression")

	// Simplest possible anonymous function (does nothing, no params, no results)
	dn := func() {}

	dn()

	// An anoymous function/lambda expression to add a percentage to a float
	result := func(value float64, p float64) float64 {
		return (value / 100) * (p + 100)
	}(50, 17.5)

	fmt.Println(result)

	// An anonymous function/lambda expression for raising a number to a given power
	pw := func(x, p int) int {

		r := x

		for i := 1; i < p; i++ {
			r = r * x
		}

		return r
	}

	printType(pw)

	fmt.Println(pw(4, 2))

	tutorial.Section("Closures")

	// A closure for determing a square of a number
	// this is a closure because it relies on, and encloses, in-scope objects - in this case the pw function
	sq := func(x int) int {
		return pw(x, 2)
	}

	until(3, sq)

	// Another closure for determing a cube of a number
	cube := func(x int) int {
		return pw(x, 3)
	}

	until(3, cube)
}

func until(max int, f func(int) int) {

	for i := 1; i <= max; i++ {
		fmt.Println(f(i))
	}

	fmt.Println()
}

func printType(i interface{}) {
	fmt.Printf("Type is %T\n", i)
}
