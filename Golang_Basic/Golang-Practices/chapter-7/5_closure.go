package main

import "fmt"

func counter() func() int {
	count := 0

	// This closure captures the 'count' variable.
	increment := func() int {
		count++
		return count
	}

	return increment
}
func main() {
	// Create a counter function
	counterFunc := counter()

	// Use the counter function to increment and retrieve values
	fmt.Println(counterFunc()) // 1
	fmt.Println(counterFunc()) // 2
	fmt.Println(counterFunc()) // 3
}
