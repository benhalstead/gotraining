package main

import (
	"fmt"
	s "github.com/benhalstead/gotraining/structures/sample"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	tutorial.Section("Struct variables")

	//Declaring an variable with a struct type gives you a valid but empty struct of that type
	var ecd s.ContactDetails

	fmt.Printf("New struct %#v\n", ecd)

	//You can initialise the contents of the struct as a literal

	cd := s.ContactDetails{
		WorkLandline: "+44123123",
		WorkMobile:   "+44456456",
	}

	fmt.Printf("Populated struct %#v\n", cd)

	tutorial.Section("Struct from new")

	// You use the builtin function new to create a new struct

	ncd := new(s.ContactDetails)

	// It is very important to note that this will give you a pointer to a struct...
	fmt.Printf("Type from new() %T\n", ncd)

	// ...but Go will allow you to access the members and methods on that struct as if it were an actual value
	// without explicit de-referencing
	ncd.WorkMobile = "+121254556"

	fmt.Printf("Struct pointer %#v\n", ncd)

	tutorial.Section("Define you own new methods")

	// If your struct requires initialisation (e.g. unexported fields need to be populated), the convention is
	// to make a function called NewTypeName in the same package as the struct
	icd := s.NewContactDetails("+44123123", "+441238432", "+44987123")

	fmt.Printf("Struct pointer %#v\n", icd)
}
