package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"time"
)

func main() {

	defer fmt.Println("main goroutine ends")

	/*
		The key concept in Go concurrency is the goroutine. A goroutine refers to a path of execution
		that can be run concurrently with other paths of execution.

		goroutines are often referred to as lightweight threads or 'green' threads, because they
		are faster to start and consume fewer resources than OS/kernel threads. Because the word
		thread is overloaded, the term goroutine is used exclusively when talking about Go.

		If you are interested in some of the concepts behind green threads and the relationship with
		kernel threads, these articles are worth reading (but entirely optional)

		https://en.wikipedia.org/wiki/Green_threads
		https://medium.com/@riteeksrivastava/a-complete-journey-with-goroutines-8472630c7f5c
		https://stackoverflow.com/questions/48638663/what-is-relationship-between-goroutine-and-thread-in-kernel-and-user-state

	*/

	tutorial.Section("Simple goroutine")

	// A new goroutine is created by putting the keyword go in front of any function call
	go waitMs(5, "Done")
	fmt.Println("After goroutine")

	// Note that these time.Sleep calls are needed in this illustration to stop the examples being run out of order!
	time.Sleep(time.Millisecond * 100)

	// You cannot easily capture the return value of the function that is being run in a goroutine
	// This won't compile go b := echoBool(true) and nor would b := go echoBool(true)

	// You can run any function in a goroutine, include a function, but be careful as the closure will have
	// access to everything that is in scope of the parent function
	tutorial.Section("Running a closure in a goroutine")

	a := 1

	go func() {
		a = 2
	}()

	time.Sleep(time.Millisecond * 100)
	fmt.Printf("I think a is set to 1, but it is %d\n\n", a)

	// goroutines can leak - it is possible to create goroutines that don't exit. This is the most
	// common way of introducing a memory leak into a Go program.

	// goroutines are independent - there is no concept of parent or child goroutines. A goroutine initiated by another
	// goroutine does not exit when the calling goroutine does
	//
	// The only exception to this is the 'main' goroutine - when that ends, all other go routines die

	tutorial.Section("goroutines can outlive calling goroutines")

	go func() {

		go tickForTwoSeconds()
		fmt.Println("Closure goroutine ends")

	}()

	time.Sleep(time.Second * 3)

	// There is no way of explicitly controlling or gather data on an individual goroutine once it has
	// been spawned. Communication and data flow between goroutines are accomplished with channels and contexts, both of
	// which are covered in later lessons
}

func tickForTwoSeconds() {

	for i := 0; i < 10; i++ {
		fmt.Println("Tick")
		time.Sleep(time.Millisecond * 200)
	}

	fmt.Println("tick goroutine ends")

}

func echoBool(b bool) bool {
	return b
}

func waitMs(i int, s string) bool {
	time.Sleep(time.Millisecond * time.Duration(i))

	fmt.Println(s)

	return true
}
