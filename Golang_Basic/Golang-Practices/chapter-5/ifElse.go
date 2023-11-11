package main

import "fmt"

func main() {
	fmt.Println("Enter a number : ")
	var num int64
	fmt.Scanf("%d ", &num)

	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

}
