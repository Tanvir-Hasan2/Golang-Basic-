package main

import "fmt"

func main() {
	var i, j = 10, 20
	fmt.Print(i, j, "\n") //insert a space between arguments neither are strings
	var firstName, lastName string = "Anik", "Adnan"
	fmt.Print(firstName, lastName, "\n") // no white space
	fmt.Println(firstName, lastName)

	fullName := fmt.Sprintf(
		"My full name is %s %s",
		firstName,
		lastName,
	)
	fmt.Println(fullName)
}
