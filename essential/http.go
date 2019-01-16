package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"net/http"
)

func main() {

	// Go's HTTP client library is similar to most modern languages. This file gives an example of a simple GET and POST

	// https://golang.org/pkg/http
	basicGet()
	basicPost()
	requestWithMoreControl()

}

func basicGet() {

	tutorial.Section("Basic GET")

	// If you need to make a simple GET request and don't need any control over headers or cookies, there is a simple helper
	// method

	if res, err := http.Get("http://ip.jsontest.com/"); err != nil {

		fmt.Printf("Error type: %T Message: %s\n", err, err.Error())

	} else {

		//http.Reponse is a struct

		//Status code is in...

		fmt.Printf("Status: %d\n", res.StatusCode)

		//Body is res.Body - it's a resource that needs to be closed
		defer res.Body.Close()

		//It implements Reader, so can be passed directly to JSON decoders
		// (see json.go in this folder for more examples)

		dc := json.NewDecoder(res.Body)
		target := make(map[string]interface{})
		dc.Decode(&target)

		fmt.Printf("%v\n", target)

	}

}

func basicPost() {

	tutorial.Section("Basic POST")

	// Create an object to use as the POST body
	m := make(map[string]interface{})

	m["A"] = 1
	m["B"] = []string{"BEE"}

	//Serialise to JSON
	j, _ := json.Marshal(m)

	//Wrap in a Reader
	r := bytes.NewReader(j)

	if res, err := http.Post("http://validate.jsontest.com/", "application/json", r); err != nil {
		fmt.Printf("Error type: %T Message: %s\n", err, err.Error())
	} else {
		fmt.Printf("Status: %d\n", res.StatusCode)
	}

}

func requestWithMoreControl() {

	tutorial.Section("Request with more control")

	// For more control over a HTTP request, you need to create a Client then a Request

	client := http.Client{
		//Control things like redirect handling and cookie support here
	}

	// This call would error if URI unparseable
	// If request needs a body, supply as a Reader to this call
	req, _ := http.NewRequest("HEAD", "http://validate.jsontest.com/", nil)

	// Once you have a request object, you can set headers
	req.Header.Add("A", "B")

	if res, err := client.Do(req); err != nil {
		fmt.Printf("Error type: %T Message: %s\n", err, err.Error())
	} else {
		fmt.Printf("Status: %d\n", res.StatusCode)
	}

}
