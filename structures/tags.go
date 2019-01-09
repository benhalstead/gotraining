package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	First  string `json:"firstname" xml:"first-name"`
	Middle string `json:"middle" xml:"middle-name"`
	Last   string `json:"firstname" xml:"first-name"`
	Age    int    `json:"age,string"`
}

type Custom struct {
	MyField string `MyTag:value`
}

func main() {

	//Tags are static meta-data attached to a struct. The most common usage is to give hints to framework code,
	// code that maps from one data structure to another. The example in Person
	// above is real tags that instruct Go's JSON and XML decoders to use different names and types when serialising and deserialsing
	//to XML and JSON

	// You can access the values of tags at runtime, but you need to use reflection
	p := Person{}

	//First step is to get the type meta-data for the struct
	t := reflect.TypeOf(p)

	//Then get the field we're interested in:
	f, _ := t.FieldByName("Age")

	//Then see if the tag we're interested in is there
	if v, okay := f.Tag.Lookup("json"); okay {
		fmt.Println(v)

	}

}
