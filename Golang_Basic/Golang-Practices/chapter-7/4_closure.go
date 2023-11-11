package main

import "fmt"

func main() {
	// Outer function
	outerVar := 10

	// Define a closure
	closure := func(x int) int {
		return x + outerVar
	}

	// Call the closure
	result := closure(5)

	fmt.Printf("Result: %d\n", result)
}
