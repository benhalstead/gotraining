package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	exampleDataRace()
	exampleMutex()
	exampleAtomic()

}

func exampleDataRace() {
	tutorial.Section("Shared resources and goroutines")

	// If two or more goroutines share access to a data structure, it is important to understand whether
	// the way the goroutines access that data is safe.
	//
	// Some data structures, like maps, are explicitly called out in the Godoc as being 'not goroutine-safe'. Sometimes
	// individual structs are generally goroutine-safe, but individual methods are not.

	// This example shows two goroutines (in closures) sharing access to an array, which we're choosing to consider as
	// a set for the purposes of this example (e.g we're worried about the contents not the order)

	v := []int{}

	oddDone := make(chan bool)
	evenDone := make(chan bool)

	loop := func(id string, c chan bool, start int) {

		for i := start; i < 10; i = i + 2 {
			v = append(v, i)
			fmt.Printf("%s %v\n", id, v)
			sleepMs(45)
		}

		c <- true
	}

	go loop("Odd", oddDone, 1)
	go loop("Even", evenDone, 0)

	<-oddDone
	<-evenDone

	fmt.Printf("Contents: %v Length: %d\n", v, len(v))

	// Sometimes when you run this code, only half of the values will be present in the array, This an example of a 'race condition' or
	// 'data race' and the effect is to make the behaviour of this code 'nondeterministic'

}

func exampleMutex() {
	tutorial.Section("Mutexes")

	// As golang does not include a 'synchronized' keyword or similar, it is the responsible of application developers
	// to understand which code is safe and which isn't and take appropriate action

	// The most common way of ensuring that 2+ goroutines access shared resources safely is to use a data structure called
	// a mutex (MUTually EXclusive)

	// A mutex is a lock or 'conch'. Only one goroutine may hold the lock at a time and all other goroutines must
	// wait for the holder to release the lock

	var mx = &sync.Mutex{}

	v := []int{}

	oddDone := make(chan bool)
	evenDone := make(chan bool)

	loop := func(id string, c chan bool, start int) {

		for i := start; i < 10; i = i + 2 {
			mx.Lock()
			v = append(v, i)
			fmt.Printf("%s %v\n", id, v)
			mx.Unlock()
			sleepMs(45)
		}

		c <- true
	}

	go loop("Odd", oddDone, 1)
	go loop("Even", evenDone, 0)

	<-oddDone
	<-evenDone

	fmt.Printf("Contents: %v Length: %d\n", v, len(v))
}

func exampleAtomic() {
	tutorial.Section("Atomics")

	// There are some circumstances where the data structure shared between goroutines is a simple numeric type
	// such as keeping track of the number of requests currently running in a web service.
	//
	// For these use case there is a package https://golang.org/pkg/sync/atomic/
	//
	// As the documentation for package emphasises, these functions and data structures are only for certain low level actions
	// and are not really suitable as mechanism for coordinating between goroutines

	// This example simulates a pair of request handlers and a counter showing how many requests are currently active

	var activeRequests int64

	getDone := make(chan bool)
	postDone := make(chan bool)

	loop := func(c chan bool) {

		for i := 1; i < 10; i++ {
			simulateRequest(&activeRequests)
		}

		c <- true

	}

	go loop(getDone)
	go loop(postDone)

	<-getDone
	<-postDone

	fmt.Printf("End count should always be zero, is: %d", activeRequests)

}

func simulateRequest(activeRequests *int64) {
	c := atomic.AddInt64(activeRequests, 1) //
	defer atomic.AddInt64(activeRequests, -1)

	fmt.Printf("Currently active: %d\n", c)

	sleepMs(25)
}

func sleepMs(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}
