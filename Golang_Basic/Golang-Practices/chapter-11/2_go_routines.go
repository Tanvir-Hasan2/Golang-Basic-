package main

import "fmt"

func main() {
	// Create a channel.
	ch := make(chan int)

	// Start a goroutine to send a value to the channel.
	go func() {
		ch <- 1
	}()

	// Receive the value from the channel and print it to the console.
	fmt.Println(<-ch)

	fmt.Println("finished main func")
}
