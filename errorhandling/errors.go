package main

import (
	"errors"
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	// Go explicitly rejects the try/catch/throw pattern for exceptions in other languages
	// https://golang.org/doc/faq#exceptions

	// Instead, they prefered approach is for functions to return errors

	// An error is any type that implements the Error interface, e.g. a type that has a method
	//
	// Error() string

	tutorial.Section("Generic error with static message")

	// The package errors provide a function to help you create a simple error that contains nothing but a message

	var e error

	e = errors.New("Simple message")

	fmt.Printf("Type: %T Error: %s\n", e, e.Error())

	tutorial.Section("Generic error with templated message")

	// The package fmt provides a function for creating a generic error using the same templating pattern as fmt.Printf

	ec := 404
	e = fmt.Errorf("HTTP error: %d", ec)

	fmt.Printf("Type: %T Error: %s\n", e, e.Error())

	tutorial.Section("Errors from functions")

	// A very common pattern is for a function to return the value you asked for and an error
	// If the error is nil, you can assume the value is safe to use
	// If the error is not nil, the value is not safe to use (it's most likely a zero value) and you should handle the error

	if v, err := valueAndError(true); err == nil {
		// This block executes if the error is nil and the value is okay
		fmt.Println(v)
	} else {
		// This block executes if the error is not nil
	}

	// That pattern is clean, but restricts the scope of the returned value to inside the if/else block
	// If you need to use the variable outside of the block, use this pattern:

	var err error
	var v int

	if v, err = valueAndError(false); err != nil {
		//This block only executes if the error is not nil - note the lack of an assignment operator ( := ) above
	}

	fmt.Printf("Value is %d\n", v)

	tutorial.Section("Custom error types")

	// Some functions do more than just use the basic errors created with errors.New and fmt.Errorf. They have
	// custom error types which may contain more detailed information. The method signature will generally
	// return the error interface, but the GoDoc will advertise which concrete types might be returned as errors

	err = connectLocal()

	fmt.Println(err.Error())

	// We can use a type assertion to check to see if the error is actually a more complex type (use a type switch if there are multiple possiblities)
	if ce, okay := err.(ConnectionError); okay {

		fmt.Printf("Blacklisting port %d\n", ce.Port)
	}

}

type ConnectionError struct {
	Server string
	Port   int
}

func (de ConnectionError) Error() string {
	return fmt.Sprintf("Could not connect to server %s on port %d\n")
}

func valueAndError(causeError bool) (int, error) {
	if causeError {
		return 0, errors.New("You asked for an error")
	} else {
		return 1, nil
	}
}

func connectLocal() error {

	return ConnectionError{
		Server: "localhost",
		Port:   8080,
	}

}
