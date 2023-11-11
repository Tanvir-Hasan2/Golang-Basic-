package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
