package main

import "fmt"

func main() {
	var num int64
	fmt.Println("enter a number : ")
	fmt.Scan(&num)
	switch {
	case num > 0:
		fmt.Println("number is positive")
	case num < 0:
		fmt.Println("number is negative")
	default:
		fmt.Println("number is zero")
	}
}
