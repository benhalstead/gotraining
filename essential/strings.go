package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"strings"
)

func main() {

	// The strings package provides functions for manipulating strings

	// https://golang.org/pkg/strings/

	// These are generally the standard set of manipulation functions you'd find in any language (index of, split)

	base := "My, test, string! 😐"

	s := strings.ToUpper(base)
	tutorial.TypeValue(s)

	s = strings.ToLower(base)
	tutorial.TypeValue(s)

	sa := strings.Split(base, ",")
	tutorial.TypeValue(sa)

	i := strings.Index(base, "😐")
	tutorial.TypeValue(i)

	i = strings.LastIndex(base, ",")
	tutorial.TypeValue(i)

	// Substring functions are NOT part of this package. This is because strings are slices and the standard slice manipulation syntax should be used instead:

	s = base[:4]
	fmt.Println(s)

	s = base[1:4]
	fmt.Println(s)

	s = base[8:]
	fmt.Println(s)

}
