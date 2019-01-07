package main

import (
	"fmt"
	"time"
)

func main() {

	// The built-in package fmt contains string formatting functions heavily based on those found in C
	// Many of those functions output directly to stdout, so are commonly used for debugging code (but not for logging,
	// there is a separate package for that

	// For basic static messages, you can use fmt.Println to print a string with a newline at the end or fmt.Print without a newline
	fmt.Println("A simple message")

	// If you want more control over the output or need to include a variable in the output, use Printf
	// Note that Printf doesn't include a newline
	name := "World"
	fmt.Printf("Hello, %s\n", name)

	// %s is called a verb and the full list of supported verbs is here https://golang.org/pkg/fmt/
	// Some common ones:

	// %s is a string
	// %d is a base 10 integer (d stands for decimal)
	// %v the value of any variable (formatting is Go's default for that type
	// %t a boolean (t)rue or false

	fmt.Printf("String: %s Int: %d Value: %v Bool: %t\n", "some string", 14, time.Now(), 1 == 2)

	// Some tips and tricks:

	//Sprintf works like Printf but returns a string instead of printing to stdout
	s := fmt.Sprintf("Test message number %d", 1)
	fmt.Println(s)

	//Errorf wraps the formatted string in an Error interface
	e := fmt.Errorf("A formatted error message from %s", "Go tutorial")
	fmt.Println(e.Error())
}
