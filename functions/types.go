package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

// A function type that can be used as a type for variables
type WriterFunc func([]string) error

// Example of a struct that has a function as a member variable
type Document struct {
	Contents []string
	Writer   WriterFunc
}

// A method which makes use of the function in the member variable
func (d *Document) Write() error {

	return d.Writer(d.Contents)
}

// A function which can be used as a WriterFunc
func FileSystemWriter(s []string) error {

	fmt.Println("Writing using FileSystemWriter")

	return nil
}

// Simple struct illustrating that methods attached to a struct can be assigned to a variable
type S3Writer struct {
	creds string
}

func (sw *S3Writer) Write(s []string) error {

	fmt.Printf("Writing using S3Writer with creds %s\n", sw.creds)

	return nil
}

// A function type with no args or return values
type close func()

// A function that returns a typed function as one of its return values
func OpenResource() (interface{}, close) {

	return "RESOURCE", func() {
		fmt.Println("CLOSE")
	}

}

func main() {

	tutorial.Section("Function type variables")

	// Variables can be declared with a function as their type
	var wf WriterFunc

	// When declared, their value is nil
	fmt.Printf("Type: %T Value: %v\n", wf, wf)

	// Any function that has exactly the same signature (even a method that is attached to a type) can be assigned to the variable
	// Note there are no brackets after the function name
	wf = FileSystemWriter

	fmt.Printf("Type: %T Value: %v\n", wf, wf)

	// Even a method that is a method on another type
	sw := S3Writer{
		"ALSKMASFd1111",
	}

	wf = sw.Write

	fmt.Printf("Type: %T Value: %v\n", wf, wf)

	tutorial.Section("Functions as variables")

	// Structs can have functions as member variables - this is a way of modifying the behaviour of a type without
	// having to create a new type (composition in OO terms)
	var d Document

	d.Writer = FileSystemWriter

	d.Write()

	d.Writer = sw.Write

	d.Write()

	tutorial.Section("Assign closure to variable")

	// As long as the method signature is correct, an anonymous function or closure can be assigned to a variable
	message := "Writing using closure"

	d.Writer = func([]string) error {

		fmt.Println(message)

		return nil
	}

	fmt.Printf("Type: %T Value: %v\n", d.Writer, d.Writer)

	d.Write()

}
