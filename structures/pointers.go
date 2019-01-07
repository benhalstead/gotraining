package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	// In Go, when you pass a variable to a function (with the exception of maps, slices and channels), the function will receive
	// a COPY of that variable - if you change it inside the function, the original value is unaffected.

	// This mechanism is known in most programming languages as call by value

	tutorial.Section("Call by value")

	a := 2
	increment(a)

	fmt.Printf("a is still %d\n", a)

	tutorial.Section("Call by value - exception for maps")

	m := map[string]int{
		"ORIGINAL": 0,
	}

	fmt.Printf("Before: %v\n", m)

	addToMap(m)

	fmt.Printf("After: %v\n", m)

	tutorial.Section("Pointers")
	// If we want a function to be able to modify (mutate) the passed variable, we need to pass a pointer to that variable
	// Formally a pointer is the address in memory of a value

	// We can obtain a pointer to any value using the ampersand symbol
	ap := &a

	// Conveniently (but sometimes confusingly) man
	fmt.Printf("ap has value %v and is type %T\n", ap, ap)

	tutorial.Section("Pointer types")

	// Pointers are typed. A pointer to an int is considered a different type from a pointer to a string
	s := "Test string"
	sp := &s

	fmt.Printf("ap is type %T and sp is type %T\n", ap, sp)

	tutorial.Section("Dereferencing")

	// You use the * symbol to get the underlying value from a pointer - this action is known as dereferencing
	var x int

	x = *ap

	fmt.Printf("x has value %d and is type %T\n", x, x)

	// Note that dereferencing a pointer while assigning it to a variable gets you a COPY of the value it points to. In this example,
	// if we change the value of x, we're not changing the value of  a

	x = 3

	fmt.Printf("a is still: %d, but x is now: %d\n", a, x)

	tutorial.Section("Call by reference")

	incrementRef(&a)

	fmt.Printf("a is now %d\n", a)

	incrementRef(&a)

	fmt.Printf("a is now %d\n", a)

}

func increment(i int) {
	i++
}

func incrementRef(i *int) {
	*i++
}

func addToMap(m map[string]int) {
	m["NEW"] = 1
}
