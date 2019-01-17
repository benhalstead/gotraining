package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	// Go's name for an ordered, fixed size collection of values is slice. These are subtly different to arrays (which also exist in Go)
	// and some Go fans get upset if you interchange the terms, but for 90% of use cases, slices behave like arrays

	// The actual difference is explained in detail here:
	// https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html

	tutorial.Section("Creating slices")

	//Slices can be declared like any other variable

	var sl []string

	fmt.Printf("Length: %d\n", len(sl))

	//Or declared as a literal
	var lit = []int{1, 2, 3}

	fmt.Printf("Length: %d\n", len(lit))

	//Or initialised using make
	ms := make([]int, 3)
	ms[0] = 1
	ms[1] = 2
	ms[2] = 3

	fmt.Printf("Length: %d\n", len(ms))

	tutorial.Section("Modifying slices")

	//You can add an item (or several items) to a slice using the append builtin (remeber to assign the result of the function to a variable)
	ms = append(ms, 4, 5, 6)

	fmt.Printf("%v Length: %d\n", ms, len(ms))

	tutorial.Section("Portions of slices")

	//You can create a new slice only containing a subset of the parent slice using this syntax

	ss := ms[1:] // Every element after and including element at index 1
	fmt.Printf("%v Length: %d\n", ss, len(ss))

	ss = ms[:3] // Every element up to but excluding element at index 3
	fmt.Printf("%v Length: %d\n", ss, len(ss))

	ss = ms[3:5] // Every element from (and including) index 3 to (and excluding) index 5
	fmt.Printf("%v Length: %d\n", ss, len(ss))

	tutorial.Section("Iterating over slices")

	//You can use the range keyword in a for loop to iterate over a slice - both the current
	//index and value are available
	for i, v := range ms {
		fmt.Printf("Value at index %d is %d\n", i, v)
	}

	//In the majority of use cases, you are not interested in the index and will receive a compiler warning if you use the
	//above pattern. You can ignore the index by using the blank identifier _ in the for declaration

	for _, v := range ms {
		fmt.Printf("%d\n", v)
	}

	tutorial.Section("Arrays")

	//If you specify a size when initialising a slice, it is an array
	var a [2]string

	fmt.Printf("Length: %d\n", len(a))

	// You can generally treat an array as if it were a slice BUT the behaviour when you pass an array to a function is different - arrays are
	// call by value, slices are call by reference

}
