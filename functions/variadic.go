package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"time"
)

// Only one argument in a function definition can accept variable numbers of arguments. It must be the last (or only)
// parameter defined.
// The ... (ellipsis) operator signifies that the parameter accepts varying numbers of values
func TypedVariadic(a string, b ...int) {

	fmt.Printf("Param b type is %T\n", b)

}

func EmptyInterfaceVariadic(i ...interface{}) {

	for _, p := range i {
		fmt.Printf("Recevied a %T\n", p)
	}
}

func recursive(b ...int) {
	recursive(b...) //Not the ellipsis AFTER the param
}

func main() {

	// Variadic functions (varargs or vargs in other languages) allow a single named parameter on a function to accept 0, 1 or n values

	tutorial.Section("Single type")

	TypedVariadic("ZERO")
	TypedVariadic("ONE", 1)
	TypedVariadic("TWO", 1, 2)

	// As the args are packed into a typed array, you can pass an slice or array of that type into a variadic function
	// using the ellipsis operator AFTER the parameter
	tutorial.Section("Pass arrays")
	ia := []int{4, 5, 6, 7, 8}

	TypedVariadic("ARRAY", ia...)

	//This technique is also useful for recursive variadic functions

	tutorial.Section("interface{} variadic")

	// If your variable parameter is of type interface{}, you can pass anything you like into the function

	EmptyInterfaceVariadic(2, "TEST", time.Now())

}
