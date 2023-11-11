package main

import "fmt"

func main() {
	fmt.Println(myFunction(1, 2))
	res, math_type := myFunction(5, 14) // store return value
	fmt.Println(res, math_type)
	result, _ := myFunction(22, 45)
	fmt.Println(result)
}

//multiple return value with named
func myFunction(x int, y int) (result int, name string) {
	result = x + y
	name = ": summation result"
	// return result
	return
}
