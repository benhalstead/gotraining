package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
)

func main() {

	tutorial.Section("Maps")

	//Maps map a key to a value. Any type can be stored as a value, 'comparable' types can be used as keys
	//See https://blog.golang.org/go-maps-in-action

	//Maps can be declared like any other type
	var sim map[string]int

	//But this will be nil until you initialise it with the builtin function make
	sim = make(map[string]int)

	//You can then write to the map...
	sim["ONE"] = 1
	sim["TWO"] = 2

	//...and read from it
	fmt.Printf("%d %d %d", sim["ONE"], sim["TWO"], sim["THREE"])

	//If you request a key that is not present in the map, you will receive the ZERO VALUE for the map's value type
	//In the example above, requesting sim["THREE"] returns 0, because that is the zero value for int

	// If you need to know whether a map actually contained the key you asked for, you use this pattern:

	if value, contains := sim["THREE"]; contains {
		fmt.Printf("Map has key. Value was %d\n", value)
	} else {
		fmt.Printf("Map did not contain key\n")
	}

	// You could replace value with _ (blank identifier) in the above example if you just wanted to check that the map
	// contained the value.

	//Maps can also be declared as literals - note the trailing comma after the last item - this is required
	tutorial.Section("Maps literals")

	lm := map[string]int{
		"ONE": 1,
		"TWO": 2,
	}

	fmt.Printf("%v\n", lm)

	tutorial.Section("Iterating")

	// You iterate over a map with the for range construct - both the current key and value are available to your code
	for k, v := range lm {
		fmt.Printf("%s = %d\n", k, v)
	}

	// Maps do not preserve order (e.g. when you iterate over a map, the keys aren't guaranteed to be returned in the same order in which you added
	// them to the map

	// You CAN modify a map, while you are iterating over it (but be sure you know what you're doing)

	// Maps are not goroutine safe - behaviour if two goroutines modify a map simultaneously is undefined. Use a mutex to lock access
	// https://blog.golang.org/go-maps-in-action

}
