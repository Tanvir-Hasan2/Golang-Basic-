package main

import (
	"fmt"
	"time"
)

func Server1(ch1 chan string) {
	time.Sleep(2 * time.Second)
	ch1 <- "server 1"
}
func Server2(ch2 chan string) {
	time.Sleep(3 * time.Second)
	ch2 <- "server 2"
}
func main() {
	// Create two channels.
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines to send values to the channels.
	go Server1(ch1)
	go Server2(ch2)

	// Multiplex between the two channels and print the received values to the console.

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}

}
