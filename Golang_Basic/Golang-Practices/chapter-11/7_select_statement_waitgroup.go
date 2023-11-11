package main

import (
	"fmt"
	"sync"
	"time"
)

func Server1(ch1 chan string, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	ch1 <- "server 1"
	wg.Done()
	close(ch1)
}
func Server2(ch2 chan string, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	ch2 <- "server 2"
	wg.Done()
	close(ch2)
}
func main() {
	// Create two channels.
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	// Start two goroutines to send values to the channels.
	go Server1(ch1, &wg)
	go Server2(ch2, &wg)

	// Wait for the other two goroutines to finish.
	wg.Wait()

	// Multiplex between the two channels and print the received values to the console.
	for {
		select {
		case val1 := <-ch1:
			fmt.Println(val1)
		case val2 := <-ch2:
			fmt.Println(val2)
		default:
			// Both channels have been closed, so exit the loop.
			close(ch1)
			close(ch2)
			break
		}
	}
	//error not fixed

}
