package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {

}

func privateFunction() {
	//Functions with a lowercase character as the first character in the function name are private to the current package
}

func PublicFunction() {
	//Functions with an uppercase character as the first character in the function name are public and can be imported and
	//used in other packages
}

func Params(a int, b string) {
	//Parameter names come before the type
}

func SameParams(a, b, c int) {
	// Only need to specify the type once in functions
}

func SingleReturn() bool {
	//Return types come at the end of the function signature
	return true
}

func MultiReturn() (bool, error) {

	//You can return n values from a function
	return false, nil
}

func NamedReturn() (b bool, err error) {
	// You can name your return values, but this is only really done to add clarity to GoDoc or to save a variable declaration
	// and less commonly to allow defer statements to alter return values (see defer.go in this folder)
	b = true
	err = errors.New("Error message")

	// You still need to explicitly return
	return b, err
}

func CallingMultiReturnValueFunctions() {

	result, err := NamedReturn()

	fmt.Printf("%v %v\n", result, err)

	//You can ignore return values with the 'blank identifier' _

	result, _ = MultiReturn()

	// IF YOU ARE IGNORING AN ERROR, YOU ARE DOING SOMETHING WRONG

}
