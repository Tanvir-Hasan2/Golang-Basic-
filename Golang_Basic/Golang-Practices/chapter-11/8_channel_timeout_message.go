package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel.
	ch := make(chan int)

	// Start a goroutine to send a value to the channel.
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	// Wait for a value to be received from the channel within 2 seconds.
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
