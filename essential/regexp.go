package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"regexp"
)

func main() {

	// Go's regex syntax is similar to PHP and Perl (and I think later versions of Ruby)
	// It is documented in great depth https://golang.org/pkg/regexp/ so this file
	// just gives some quick examples about matching and group capture

	// This is a good site for developing and testing your regex https://regex-golang.appspot.com/assets/html/index.html

	tutorial.Section("Basic matching")
	// Patterns must be 'compiled' before they can be used

	currencyPattern := "^\\d*\\.\\d{2}$"
	//Defining patterns in code is ugly because backslashes need to be escaped this can be read as
	// ^\d*\.\d{2}$   n digits, followed by a period, followed by exactly two digits (a currency value)

	// If your patterns can change at runtime, use:

	if crx, err := regexp.Compile(currencyPattern); err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Printf("Match: %t\n", crx.MatchString("100.00"))

	}

	// If your patterns are hardcoded, you can use:

	crx := regexp.MustCompile(currencyPattern)

	// Which will panic if the regex is illegal

	fmt.Printf("Match: %t\n", crx.MatchString("10!00"))

	tutorial.Section("Capture groups")

	// Like most regex libraries, Go allows you to capture portions of a match by surrounding the group your're interested in brackets
	currencyPattern = "^(\\d*)\\.(\\d{2})$"
	crx = regexp.MustCompile(currencyPattern)

	groups := crx.FindStringSubmatch("100.54")

	// In the resulting array, the element at zero is the portion of the string that matches the overall regex
	// The first group is at index 1

	pounds := groups[1]
	pence := groups[2]

	fmt.Printf("Pounds: %s Pence: %s\n", pounds, pence)

}
