package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i += 1

	}
	fmt.Println()

	for j := 2; j <= 100; j += 10 {
		fmt.Println(j)
	}
}
