package channel

import "fmt"

// Pump test
func Pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// Suck test
func Suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
