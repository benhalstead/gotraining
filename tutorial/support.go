package tutorial

import (
	"fmt"
	"time"
)

func Section(m string) {
	fmt.Printf("\n\n%s:\n\n", m)
}

func TypeValue(i interface{}) {
	fmt.Printf("Type: %T Value: %v\n", i, i)
}

func Tick(count, ms int) {

	for i := 0; i < count; i++ {
		fmt.Println("Tick")
		time.Sleep(time.Millisecond * time.Duration(ms))
	}

}
