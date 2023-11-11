package main

import "fmt"

func main() {
	//anonymous function
	add := func(x, y int64) int64 {
		return x + y
	}
	fmt.Println(add(5, 8))

	n := 3
	increment := func() int {
		n++
		return n
	}
	fmt.Println(increment()) // 3
	fmt.Println(increment()) // 4

}
