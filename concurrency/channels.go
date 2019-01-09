package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"net/http"
	"time"
)

func main() {

	tutorial.Section("Initialising channels")

	// Channels are a type and they can be declared as other variables are...
	// The type after the word chan is the data-type
	var c chan int

	tutorial.TypeValue(c)

	// But, like maps, they are nil until you initialise them with make, so it more common to see
	// them declared and initialised in one statement:

	ic := make(chan int)

	// By default, channels are blocking - if you write to the channel you cannot write to it again until
	// it has been read from

	tutorial.Section("Reading from and writing to blocking (unbuffered) channels")

	// Channels don't make sense in an application with only one (main) goroutine, so these examples will use
	// anonymous functions to spawn other goroutines

	f := func(fc chan int) {
		for i := 0; i < 2; i++ {
			//Write an int to the channel
			fmt.Printf("About to write %d\n", i)
			fc <- i
			fmt.Printf("Wrote %d\n", i)

		}
	}

	go f(ic)

	// Read values from the channel (you would normally do this in a loop of some sort)

	// Sleep to illustrate that the write operation blocks until the value is read from the channel
	tutorial.Tick(5, 200)

	a := <-ic

	b := <-ic

	fmt.Printf("%d %d\n", a, b)

	tutorial.Section("Reading from and writing to buffered channels")

	ic = make(chan int, 2)

	// If you call make on a channel and provide an int argument, that is the number of items this channel will
	// allow to be written to it without a read before it blocks

	go f(ic)

	// Sleep to illustrate that the write operation is not blocked (because the channel has a buffer of 2)
	tutorial.Tick(5, 200)

	a = <-ic

	b = <-ic

	fmt.Printf("%d %d\n", a, b)

	tutorial.Section("Types")

	// You can declare a channel of any type (including other channels)
	// Be careful when sending pointers and maps as you lose control over which goroutine is accessing the underlying
	// data. Channels are safest when sending values.

	type S struct{}

	cc := make(chan chan int)
	tutorial.TypeValue(cc)

	sc := make(chan S)
	tutorial.TypeValue(sc)

	spc := make(chan *S)
	tutorial.TypeValue(spc)

	mc := make(chan map[string]string)
	tutorial.TypeValue(mc)

	tutorial.Section("Closing channels")

	// The sending side of a channel can choose to close it. The receiving side can test to see if the channel has been closed
	// and can stop waiting for data to be sent

	c = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Printf("Sent %d\n", i)
		}

		close(c)
	}()

	// A receiver can range over a channel - this loop is infinite until the sender closes the channel
	for n := range c {
		fmt.Printf("Received %d\n", n)
	}

	// If you are reading from a channel outside of a range, you can use this syntax to check for a closed channel

	tutorial.Section("Select and example")

	// In this example we fetch three web pages concurrently using channels
	// and the channel specific select control structure
	// to detect when the work is done

	if n, open := <-c; open {
		fmt.Printf("Received %d\n", n)
	} else {
		fmt.Printf("Channel closed\n")
	}

	var googleResp *http.Response
	var bbcResp *http.Response
	var cfResp *http.Response

	// In this example we will create a separate 'done' channel for each page we're requesting
	// this could be done with a single channel (and would be a good refactoring exercise)
	googleDone := make(chan bool)
	bbcDone := make(chan bool)
	cfDone := make(chan bool)

	// Spawn goroutines to fetch each URL, passing a pointer to an object to store the result in
	go get("http://www.google.com", &googleResp, googleDone)
	go get("http://www.bbc.com", &bbcResp, bbcDone)
	go get("http://www.cloudfactory.com", &cfResp, cfDone)

WaitLoop:
	for {
		// The select control structure is only useful for channels
		// Each case in the select will fire if a value can be read from the channel
		select {
		case <-bbcDone:
			fmt.Printf("BBC page done\n")
		case <-googleDone:
			fmt.Printf("Google done\n")
		case <-cfDone:
			fmt.Printf("Cloudfactory done\n")
		default:
			// The default case will fire if none of the other cases do
			if googleResp != nil && bbcResp != nil && cfResp != nil {
				break WaitLoop
			}

			//In these circumstances it's common to sleep for a short time to stop the surrounding
			//for loop consuming resources doing nothing but checks This time should be proportional
			// to the activity in question
			time.Sleep(time.Millisecond * 50)
		}

	}

	fmt.Printf("Google response was %d\n", googleResp.StatusCode)
	fmt.Printf("BBC response was %d\n", bbcResp.StatusCode)
	fmt.Printf("CF response was %d\n", cfResp.StatusCode)

}

func get(url string, result **http.Response, finish chan bool) {

	if r, err := http.Get(url); err == nil {
		*result = r
	} else {
		fmt.Printf("Error reading %s: %s\n", url, err.Error())
	}

	finish <- true
}
