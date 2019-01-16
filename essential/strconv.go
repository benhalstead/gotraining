package main

import (
	"github.com/benhalstead/gotraining/tutorial"
	"strconv"
)

func main() {

	//The strconv package contains functions for converting from text representations of a value to typed representations
	// https://golang.org/pkg/strconv

	// The most common is Atoi (ascii to int)

	v, _ := strconv.Atoi("12") // You will get an error if it can't be parsed

	tutorial.TypeValue(v)

	// Atoi assumes base 10. If you have other number bases or need to constrain the created int to certain size in bits:

	x, _ := strconv.ParseInt("1f", 16, 8)

	tutorial.TypeValue(x)

	// ParseInt returns an int64, regardless of the bitsize you pass in, but specifying the bitsize means you
	// are safely able to convert the int64 to the int version you need (the parse would've failed if the value didn't fit
	// in the specified bit size

	tutorial.TypeValue(int8(x))

	// There is a separate function ParseUint for parsing unsigned integers

	// The same rules apply to parsing floats

	f, _ := strconv.ParseFloat("0.11e10", 32)

	tutorial.TypeValue(f)

	// ParseBool accepts true and false, but also some common alternative representaions of true and false

	//ParseBool returns the boolean value represented by the string. It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.

	b, _ := strconv.ParseBool("t")

	tutorial.TypeValue(b)

}
