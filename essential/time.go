package main

import (
	"fmt"
	"time"
)

func main() {

	// Date parsing is peculiar in Go. Instead of using symbols to represent a date format, you need to remember a particular date and time:
	//
	// Jan 2 15:04:05 2006 MST
	//
	// (MST is UTC -0700)

	// A 'layout' is then that date/time expressed in the format you want to capture

	layout := "2006-01-02" //This is YYYY-MM-YY because 01 in the epcoh is January

	value := "2019-04-30"

	if d, err := time.Parse(layout, value); err == nil {
		fmt.Printf("%v\n", d)
	} else {
		//Unparseable
		fmt.Println(err.Error())
	}

	// Layouts without time components will result in a date-time with a 0000 UTC time

	layout = "1 January 2006 3pm MST"

	value = "2 February 2019 9am GMT"

	if d, err := time.Parse(layout, value); err == nil {
		fmt.Printf("%v\n", d)
	} else {
		//Unparseable
		fmt.Println(err.Error())
	}

	// Intervals between two date-times are represented as a time.Duration which is an int64. The raw value is in nano seconds
	// To convert between the nano second value and a higher order unit (ms, second, minute, hour) use the constants in the time package

	start := time.Now()

	time.Sleep(2 * time.Second)

	end := time.Since(start)

	fmt.Printf("Elaspsed in nano seconds %d\n", end)

	fmt.Printf("Elapsed using fmt's default format %v\n", end)

	fmt.Printf("Elapsed in milliseconds %d\n", end/time.Millisecond)

}
