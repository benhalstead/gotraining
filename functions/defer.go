package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	// Defer is Go's 'finally' mechanism to allow ensure resources are closed or other cleanup activity to be performed
	// reliably when a function ends.

	// Defer allows you to specify a function to execute when a function ends (but before it returns)
	tutorial.Section("Single defer")
	singleDefer()

	// The defered function can be a closure
	tutorial.Section("Closure defer")
	closureDefer()

	// You can have multiple defers in a function and they are executing in reverse order (i.e. there is a 'stack' of deferred functions)
	tutorial.Section("Multi defer")
	multiDefer()

	// Be very careful using defer in loops
	tutorial.Section("Loop defer")
	loopDefer()

	// Defer runs even if code in the function panics (and is the idiomatic way to recover from panics)
	tutorial.Section("Loop defer")
	panicDefer()

}

func singleDefer() {

	openResource()
	defer closeResource()

	fmt.Println("After the defer line")

}

func closureDefer() {
	a := 2

	defer func() {
		fmt.Printf("Defer %d\n", a)
	}()

	a++
}

func multiDefer() {

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")

}

func loopDefer() {

	for i := 0; i < 3; i++ {
		openResource()
		defer closeResource()
	}

}

func panicDefer() {

	openResource()

	defer func() {
		closeResource()

		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	panic("Example panic")

}

func openResource() {
	fmt.Println("Open resource")
}

func closeResource() {
	fmt.Println("Close resource")
}
