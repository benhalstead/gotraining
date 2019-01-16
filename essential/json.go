package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"strings"
)

func main() {

	/*
		JSON unmarshaling (into an object from a JSON file or stream) and marshalling (from an object into
		a file or a stream) is built-in to Go in the json package

		There are built-in functions from mashalling and unmarshalling from byte array representations of JSON
		and Encoder and Decoder types for working with Readers and Writers (see the bufio.go file in this package for
		more info on that)

		The most common use case when working with JSON is to marshall in and out of structs, but it is almost equally
		common to use a map[string]interface{} instead - this is safer in some ways as some types of schema changes
		(new fields added) can be mitigated against.

		https://golang.org/pkg/json

	*/

	unmarshallIntoStructFromReader()
	unmarshallIntoMapFromReader()
	unmarshallFromString()
	marshallFromStructToBytes()
	marhsallFromStructToWriter()
	marhsallToPrettyPrintedString()
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

func unmarshallIntoStructFromReader() {

	tutorial.Section("Unmarshall into struct from Reader")

	// Readers are an abstraction over any data type that can be intepreted as a stream of bytes.
	// Once the input source is modelled as a Reader, the clients of the Reader don't care what the underlying
	// source is - it could be a string, a file or a network connection

	sr := strings.NewReader(simpleJSON)

	// We can create a Decoder over our Reader which will provide unmarshalling functions
	dc := json.NewDecoder(sr)

	// We need a variable to decode into

	var t Target

	if err := dc.Decode(&t); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%#v\n", t)
	}
}

func unmarshallIntoMapFromReader() {

	tutorial.Section("Unmarshall into map from Reader")

	// It can sometimes be more flexible to parse into a semi-structured map that having to maintain a struct for every API you work with

	sr := strings.NewReader(simpleJSON)
	dc := json.NewDecoder(sr)

	//JSON field names are always strings, so the map type should be map[string]interface{}
	target := make(map[string]interface{})

	if err := dc.Decode(&target); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%#v\n", target)
	}

}

func unmarshallFromString() *Target {

	tutorial.Section("Unmarshall from string")

	// Generally you will be working with Readers, but sometimes JSON might be made available to your code in a string
	// or byte array.

	// If the data is an a string, you'll need convert it to []byte

	var t Target

	j := []byte(simpleJSON)

	if err := json.Unmarshal(j, &t); err != nil {
		fmt.Println(err.Error())

		return nil

	} else {
		fmt.Printf("%#v\n", t)
		return &t
	}

}

func marshallFromStructToBytes() {
	tutorial.Section("Marshall from struct to []byte")

	// Normally you will want to marshall directly to a Writer (to send a JSON object over a network or
	// to disk), but sometimes you want to keep a representation in memory

	v := unmarshallFromString()

	if marhsalled, err := json.Marshal(v); err == nil {
		fmt.Printf("Marshalled to []byte of length %d\n", len(marhsalled))
	} else {
		fmt.Println(err.Error())
	}
}

func marhsallFromStructToWriter() {
	tutorial.Section("Marshall from struct to Writer")

	v := unmarshallFromString()

	// Buffer is a type of data structure that will hold an arbitrary length sequence of bytes in memory.
	// It implements io.Writer, but you have to remember to pass a pointer to the Buffer instead of the struct itself.
	// In most cases you'll be writing to a file or network stream instead
	var b bytes.Buffer

	en := json.NewEncoder(&b)

	if err := en.Encode(v); err != nil {
		fmt.Println(err.Error())
	} else {

		// You can 'flush' a buffer to a string
		result := b.String()

		fmt.Printf("%s", result)
	}
}

func marhsallToPrettyPrintedString() {
	tutorial.Section("Pretty print")

	v := unmarshallFromString()

	// You can use MarshalIndent to control the presentation of a JSON. This example inserts a tab in front of
	// each logically indented structure

	if j, err := json.MarshalIndent(v, "", "\t"); err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Printf("%s\n", string(j))
	}

}
