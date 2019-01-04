package main

import (
	"fmt"
	"github.com/benhalstead/gotraining/annotation"
)

func main() {

	annotation.Section("Maps")

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

	//Maps can also be declared as literals - note the trailing comma after the last item - this is required
	annotation.Section("Maps literals")

	lm := map[string]int{
		"ONE": 1,
		"TWO": 2,
	}

	fmt.Printf("%v\n", lm)

	annotation.Section("Iterating")

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
