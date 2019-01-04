package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/annotation"
)

func main() {

	// A variable can be delcared with a specific type, declared and initialised in one statement
	annotation.Section("Declarations and initialisation")

	x := 2
	var y int
	var z int = 2

	fmt.Printf("%T %T %T\n", x, y, z)

	annotation.Section("interface{} values")

	// Once the type of a variable is determined, either implicitly or explicitly, you cannot change it
	// x = "test" would not work. If you need this sort of behaviour, you must declared you variable as type interface{}

	var v interface{}

	v = 1
	fmt.Printf("%T\n", v)

	v = "1"
	fmt.Printf("%T\n", v)

	// Unlike C, a declared but not initialised variable has a predictable value according to its type
	annotation.Section("Zero values")

	var i int
	fmt.Printf("Int family (all sizes and signedness) is zero: %v\n", i)

	var b bool
	fmt.Printf("Boolean is false: %v\n", b)

	var s string
	fmt.Printf("String is empty (zero length) string: \"%v\"\n", s)

	var f float64
	fmt.Printf("Float family (both sizes) is zero: %v\n", f)

	var iface interface{}
	fmt.Printf("Interface is nil: %v\n", iface)

	var ip *int
	fmt.Printf("All pointers are nil: %v\n", ip)

	// Struct zero values are explained in the struct section

}
