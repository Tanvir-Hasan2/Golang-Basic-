package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var num float64
	fmt.Scanf("%f", &num)

	output := num * 10
	fmt.Println(output)
	output += 50
	fmt.Println(output)

}
