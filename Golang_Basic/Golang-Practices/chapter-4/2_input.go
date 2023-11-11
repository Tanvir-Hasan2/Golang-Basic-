package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	fmt.Println("Enter your name and age in the format \"name age\":")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	var fullname string
	fmt.Println("Enter your name:")
	fmt.Scanln(&fullname)

	fmt.Printf("Hello, %s!\n", fullname)

}
