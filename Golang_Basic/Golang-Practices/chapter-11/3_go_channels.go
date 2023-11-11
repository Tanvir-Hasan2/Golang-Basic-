package main

import "fmt"

func add(ch chan int, a ...int) {
	sum := 0
	for _, num := range a {
		sum += sum + num
	}
	ch <- sum
}
func main() {
	// Create a channel.
	ch := make(chan int)

	// satrt a go routine
	go add(ch, 5, 12, 25, 35)

	// value send and hold to summation
	summation := <-ch

	// Receive the value from the channel and print it to the console.
	fmt.Println(summation)

	close(ch)

	sum := <-ch // 0 assigned
	//fatal error: all goroutines are asleep - deadlock!
	// first close channel and deafult zero value assign to the var

	fmt.Println(sum)
}
