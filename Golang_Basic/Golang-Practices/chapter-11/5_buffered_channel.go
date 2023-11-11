package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 3 means it can store up to 3 values before blocking.
	ch := make(chan int, 3)

	// Start a goroutine to send values to the channel.
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// Iterate over the values of the channel and print them to the console.
	for v := range ch {
		fmt.Println(v)
	}
}
