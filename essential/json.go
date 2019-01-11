package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	/*
		JSON marshaling (into an object from a JSON file or stream) and unmarshalling (from an object into
		a file or a stream) is built-in to Go in the json package

		There are built-in functions from mashalling and unmarshalling from byte array representations of JSON
		and Encoder and Decoder types for working with Readers and Writers (see the bufio.go file in this package for
		more info on that)

		The most common use case when working with JSON is to marshall in and out of structs, but it is almost equally
		common to use a map[string]interface{} instead - this is safer in some ways as some types of schema changes
		(new fields added) can be mitigated against.

	*/

	marshallIntoStructFromReader()
}

var simpleJSON = `
{
	"numberVal": 54.1,
	"boolVal": true,
	"stringVal": "hello",
	"numArray": [1,2.0,3],
	"boolArray": [true, false],
	"stringArray": ["a","b","c"],
	"objectVal": {
		"numberVal": 5
	},
	"objectArray": [
		{
			"stringVal": "A"
		},
		{
			"stringVal": "B"
		}
	]
}
`

//  JSON from third party systems often works with camel case variable names
//  which are incompatible with Go's requirement that exported fields start with a capital
//  the use of tags (see other lessons) works around this
type Target struct {
	NumberVal   float64   `json:numberVal`
	BoolVal     bool      `json:boolVal`
	StringVal   string    `json:stringVal`
	NumArray    []float64 `json:numArray`
	BoolArray   []bool    `json:boolArray`
	StringArray []string  `json:stringArray`
	ObjectVal   *Target   `json:objectVal`
	ObjectArray []Target  `json:objectArray`
}

func marshallIntoStructFromReader() {

	// Readers are an abstraction over any data type that can be intepreted as a stream of bytes.
	// Once the input source is modelled as a Reader, the clients of the Reader don't care what the underlying
	// source is - it could be a string, a file or a network connection

	sr := strings.NewReader(simpleJSON)

	// We can create a Decoder over our Reader which will provide marshalling functions
	dc := json.NewDecoder(sr)

	// We need a variable to decode into

	var t Target

	if err := dc.Decode(&t); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%#v\n", t)
	}

}
