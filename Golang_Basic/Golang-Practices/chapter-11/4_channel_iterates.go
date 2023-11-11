package main

import "fmt"

func main() {
	// Create a channel.
	ch := make(chan int)

	// Start a goroutine to send values to the channel.
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Iterate over the values of the channel and print them to the console.
	for v := range ch {
		fmt.Println(v)
	}
}
