package main

import "fmt"

func main() {
	// var x []int64 // length of x is zero
	y := make([]int64, 4)
	y[3] = 200
	fmt.Println(y)

	a := make([]int64, 5, 10) // 3rd perameter capacity

	fmt.Println(len(a)) // underlying array size 5
}
