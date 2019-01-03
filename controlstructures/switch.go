package main

import (
	"fmt"
	ann "github.com/benhalstead/gotraining/annotation"
)

func main() {

	//You can switch on any statement that evaluates to a value:
	ann.Section("Switch on bool evaluation")

	switch 1 == 2 {
	case true:
		fmt.Println("Unlikely")
	case false:
		fmt.Println("Expected")
	default:
		//Compiler won't catch impossible to reach default statements
	}

	//Most commonly you will test for a known value
	ann.Section("Switch on known value")
	a := 3
	const SOME_CONST = 4

	switch a {
	case 3:
		fmt.Println("Expected")
	case SOME_CONST:
		fmt.Println("Expected")

	}

	// But it is possible to have a case that matches based on another variable
	ann.Section("Variable in case check ")

	a = 2
	b := 2

	switch a {
	case b:
		fmt.Printf("%d == %d\n", a, b)
	default:
		fmt.Println("Unexpected")
	}

	// There is no concept of fallthrough in Go Switch statements - one or zero cases will match and be executed
	ann.Section("No fall through")
	a = 1
	b = 1
	c := 1

	switch a {
	case b:
	case c:
		fmt.Println("Will never be reached")
	}

	//Switch statements are an idiomatic way of handling variables where you don't know the type
	ann.Section("Type switch")

	var i interface{}

	i = 1

	switch t := i.(type) {
	case int:
		fmt.Println("Found an int")
	case string:
		fmt.Println("Found a string")
	default:
		// t is i 'cast' to the actual type
		fmt.Printf("Value is %v\n", t)
	}

}

func helper() bool {
	return false
}
