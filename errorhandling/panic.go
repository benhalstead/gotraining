package main

import (
	"fmt"
)

func main() {

	// A panic is Go's mechanism for a situation that the runtime can't cope with or is generally outside of the
	// general error handling pattern of Go.
	//
	// Many panics are caused by situations analogous to those that cause runtime exceptions in other languages - nil pointer access,
	// out of bounds array access, divide by zero etc

	// A panic will bubble up through all your layers of code until it is recovered using the special keyword recover
	// recover only works when called in the context of a defer

	// If the panic reaches the top of a goroutine and that goroutine is the only one running in your application, your application will
	// exit with an error

	// As a rule of thumb, your code should never initiate a panic and your application logic should never worry about recover

	// This is generally the job of whichever framework you are using. For example, Granitic recovers from panic in
	// ws.WsHandler.ServeHttp which means it's impossible for a panic in any web service request to crash the application

	// The circumstances under which your application logic might need to protect against panics is when it's vital
	// that the flow of control is not lost. For example, a complex multi-service transaction might need to be rolled backed if
	// one part of the process causes a panic

	// The function below illustrates how to raise a panic and recover from it using a defer

	recoverable()

}

func recoverable() {

	//Common to use an anonymous function to handle recovery

	defer func() {
		// Defer will always be called, even if the code below didn't panic

		if r := recover(); r != nil {
			fmt.Printf("Recovered (%T %v)", r, r)
		}

	}()

	panicSource(0)

}

func panicSource(i int) {
	_ = 4 / i
}
