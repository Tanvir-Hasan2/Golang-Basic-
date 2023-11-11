package main

import "fmt"

func main() {
	func() {
		fmt.Println("annonumous function call")
	}()

	show := func() string {
		return "show method"
	}
	fmt.Println(show())

	display := func(a int, b int) {
		fmt.Println(a + b)
	}
	display(2, 5)
}
