package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

type Named interface {
	Name() string
}

type Measurable interface {
	Length() int
}

func NewPerson(f, l string) *Person {
	p := new(Person)
	p.Last = l
	p.First = f

	return p
}

type Person struct {
	First string
	Last  string
}

// Having this method means we implement Named
func (p Person) Name() string {
	return fmt.Sprintf("%s %s\n", p.First, p.Last)
}

// Having this method means we implement Measurable
func (p Person) Length() int {
	return len(p.First) + len(p.Last)
}

// Having this method means we implement Stringer - a builtin Go interface which means
// many Go functions will use this method to get a representation of this object
func (p Person) String() string {
	return p.Name()
}

//Any type can implement an interface
type mLength int

func (ml mLength) Length() int {
	return 1
}

// Because the check for implementation is whether you implement the methods on an interface, it is not possible
// to have no-method 'marker' interfaces in Go. The convention for simulating this is in Go is to define a single method
// on the interface with the same name as the interface.
type Marker interface {
	Marker()
}

func main() {

	// In Go, any type (not just structs) can implement an interface. An interface is considered to be implemented as long
	// as the type has exactly the same set of methods with the same method signatures as those defined by an interface

	// You can declare a variable to be an interface type
	var i Measurable

	tutorial.Section("Underlying type")

	// Any type implementing that interface can be assigned to the variable
	i = new(Person)

	// Asking for the variable's type will return the underlying type, not the interface
	fmt.Printf("i is %T\n", i)

	i = mLength(1)

	fmt.Printf("i is %T\n", i)

	tutorial.Section("Type assertion")

	// Given an interface, you can attempt to 'cast' it to an underlying type using a type assertion
	// This is the safe way if it's possible that the type is something other than you think
	if length, okay := i.(mLength); okay {
		fmt.Printf("Value of mLength %d\n", length)
	} else {
		//Recover if the type was something other than expected
		fmt.Println("Unexpected type")
	}

	// The unsafe way, if you're 100% sure you know the underlying type
	length := i.(mLength)
	fmt.Printf("Value of mLength %d\n", length)

	// If you get that wrong, a panic will be thrown

	tutorial.Section("Passing to functions")

	// As long as you type implements an interface, you can pass it to any function that expects that interface
	// You do not need to 'cast' to interface (in fact, you can't)

	var p Person
	accept(p)

	tutorial.Section("Inadvertent implementations")

	//Sometimes you do not know your code is implementing an interface! If your type has a method
	// String() string
	// It is actually implementing the fmt.Stringer interface and the behaviour of the %v verb in the printf functions will change

	p.First = "My"
	p.Last = "Name"

	fmt.Printf("Value using %%v is %v\n", p)

}

func accept(i Named) {

	fmt.Printf("accept: %s", i.Name())

}
