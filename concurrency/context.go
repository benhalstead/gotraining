package main

import (
	"context"
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	// Contexts are a late addition to Go (1.7) and are designed to solve two problems with goroutines
	//
	// 1. Once created, there is no way of telling a goroutine to terminate early
	// 2. There is no concept in Go of 'thread' or 'request' variables (variables that should be transparently pass through
	//  a function call stack so they can be used later)
	//
	// The best illustration of why 1. is a problem is a web server. A client of a web server can terminate a page load early
	// In this case they are not interested in the eventual response to their request, but the request has already been
	// received by the server and it may have already spawned goroutines to build the page
	//
	// For 2, it is very common for there to be incoming data to a web-service (often in headers) that are of no interest
	// to the logic of your application, but that need to be passed on to any downstream web-services.
	//
	// What has emerged is a convention where all of the functions in a web service request's main flow of control
	// accept a context as the first argument and pass it down to other functions
	//
	// func SomeFunction(ctx context.Context, otherArgs int)
	//
	// It is a strong convention to name the variable ctx
	//
	// https://golang.org/pkg/context/
	// https://blog.golang.org/context

	// The decision to solve two different problems with one data structure means that the documentation for contexts is unfortunately confusing

	creatingContextsExample()
	avoidingKeyCollisionsExample()

	//cancellationExample()

}

func creatingContextsExample() {

	tutorial.Section("Creating contexts (without cancel functions)")

	// Contexts exist as a hierarchy. The entry point into your framework is responsible for creating a parent (or 'root') Context
	// (for Granitic web services this happens in facility/httpserver/HttpServer.handleAll)

	// A particular method is used to create this root context

	ctx := context.Background()

	// When the context is manipulated (say, when adding a variable) a new context is DERIVED from the existing one
	// the functions context.With* are used for this
	ctx = context.WithValue(ctx, "authToken", "DFASDAS;X1]a")
	ctx = context.WithValue(ctx, "permissions", "AbbC,A11d")

	// Any values present in the parent context are copied into the derived context
	// Values can be read from the context

	at := ctx.Value("authToken")
	println(at)
	// but will be returned as type interface{} it is common to provide helper functions for reading values back out of a context (see other examples)

}

type key string

var authAsKey = key("authToken")
var authAsString = "authToken"

func avoidingKeyCollisionsExample() {
	tutorial.Section("Avoiding key collisions")

	//Contexts store data as key value pairs. There is nothing to stop code further down your request path (in your application
	//code or in third-party libraries) overwriting your value by using the same key

	// The convention for solving this problem is to make clever use of Go's type system
	// (outside this function we have declared:
	//   type key string
	// an unexported type based on string.

	ctx := context.Background()

	// Keys are defined in constants above
	ctx = context.WithValue(ctx, authAsString, "1234")
	ctx = context.WithValue(ctx, authAsKey, "ABCD")

	k := ctx.Value(authAsString)

	tutorial.TypeValue(k)

	j := ctx.Value(authAsKey)
	tutorial.TypeValue(j)

	// For this to work, your type must remain unexported (otherwise other code to use your type and overwrite your value)
	// This means that you will need to declare helper functions to read your data out of a context

	fmt.Println(ReadAuthTokenZeroVal(ctx))

	if v, okay := ReadAuthTokenOkay(ctx); okay {
		fmt.Println(v)
	}

}

// A helper function for reading a value from a context that returns a zero value if the key is not found or is not the expected type
func ReadAuthTokenZeroVal(ctx context.Context) string {

	result := ""

	if v := ctx.Value(authAsString); v != nil {

		if s, okay := v.(string); okay {
			result = s
		}

	}

	return result
}

// A helper function for reading a value from a context that returns an additional bool to indicate if a key is missing
// or not the expected type
func ReadAuthTokenOkay(ctx context.Context) (string, bool) {

	if v := ctx.Value(authAsString); v != nil {

		if s, okay := v.(string); okay {
			return s, true
		}

	}

	return "", false
}

func cancellationExample() {
	// ??
}
