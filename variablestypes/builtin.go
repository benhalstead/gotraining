package main

import (
	"fmt"
	ann "github.com/benhalstead/gotraining/annotation"
	"strconv"
)

// There a small number of built-in, basic or native types
func main() {

	//Booleans
	ann.Section("Booleans")
	a := true
	b := false

	fmt.Printf("Does a == b: %t", a == b)

	ann.Section("Parsing bools")

	// The compiler only recognises lowercase true or false, but the standard parseBool function
	// handles 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	if parsed, err := strconv.ParseBool("True"); err == nil {
		fmt.Printf("Parsed as %t\n", parsed)
	} else {
		fmt.Println("Unable to parse as a bool")
	}

	//Strings
	ann.Section("Strings")

	//Strings are immutable sequences of bytes - you may include unicode in strings
	s := "Hello 😐!"

	fmt.Println(s)

	//You can iterate and index strings as if they were arrays, but doing this will return variables of type rune
	for i, c := range s {
		fmt.Printf("Rune value %v is character %#U and starts at byte %d \n", c, c, i)
	}

	// A single 'character' can be defined in single quotes - this will be a rune (which is an alias of int32)
	c := '😐'

	fmt.Printf("%T\n", c)

	//Integers
	ann.Section("Sized integers")

	//Go has specific types for different bit lengths of int
	// int8 int16 int32 int64

	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1

	//You must cast these types to compare them
	fmt.Printf("Are they all equal? %t\n", (i8 == int8(i64)) && (int8(i16) == int8(i32)))

	//There are also unsigned variants of these four types uint8, uint16, uint32, uint64

	ann.Section("int and uint")

	// There are two other integer types int and uint
	// The size of this depends on the platform you are running on. On a 64-bit platform an int is 64-bits wide
	// BUT the compiler will not let you pass an int to a function that is expecting an int64 as the compiler
	// does not know the target platform for your code

	var pi int

	fmt.Printf("%T %T\n", pi, i64)

	//Most Go code just uses int and uint, but I have seen problems with this approach with 32 bit Windows.

	ann.Section("Floats")

	//Go provides two floating point types float32 and float64 - you must always explicitly choose a size
	//If you declare a float literal, the compiler will pick float64

	f := 1.99

	fmt.Printf("%T\n", f)

	//Floats can be converted to ints but you lose all information after the decimal place
	fmt.Printf("%d\n", int(f))

}
