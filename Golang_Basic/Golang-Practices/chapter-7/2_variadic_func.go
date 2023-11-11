package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3))
	x := []int{5, 12, 13, 16}
	fmt.Println(add(x...)) //unpacked x
}
func add(args ...int) int { //packed by args
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
