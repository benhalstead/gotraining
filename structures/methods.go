package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"strconv"
	"strings"
)

type name struct {
	First  string
	Middle string
	Last   string
}

// Functions are attached to a struct if they include a 'receiver'. A receiver is a type in brackets before the function name
// Once a function is attached to a struct in this way it is referred to as a method

//If your receiver is a pointer (*name here) it will receive a reference to a type and can therefore modify its members
func (n *name) Normalise() {
	n.First = strings.ToUpper(n.First)
	n.Middle = strings.ToUpper(n.Middle)
	n.Last = strings.ToUpper(n.Last)
}

//If your receiver is not a pointer it will receive a copy of the value
func (n name) FullName() string {
	return fmt.Sprintf("%s %s %s", n.First, n.Middle, n.Last)
}

type myInt int

func (mi myInt) ToString() string {
	return strconv.Itoa(int(mi))
}

type pseduoSet map[string]bool

func (ps pseduoSet) Add(key string) {
	ps[key] = true
}

func (ps pseduoSet) Contains(key string) bool {
	return ps[key]
}

func main() {

	tutorial.Section("Struct methods")

	// You can call methods on structs or pointers to structs

	nn := new(name)

	println(nn.FullName())

	pn := name{
		"Ben",
		"James",
		"Halstead",
	}

	fmt.Printf("%s\n", pn.FullName())

	tutorial.Section("Mutable")

	// If the method was declared with a pointer receiver, it can modify the contents of that struct
	pn.Normalise()
	fmt.Printf("%s\n", pn.FullName())

	tutorial.Section("Methods on other types")

	var mi myInt = 3

	fmt.Printf("String val: %s\n", mi.ToString())

	set := make(pseduoSet)

	set.Add("APPLE")

	fmt.Printf("Contains APPLE? %t\n", set.Contains("APPLE"))

}
