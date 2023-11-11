package main

import "fmt"

func main() {
	// if initial_statemant; condition{
	// }
	name := "Anik Adnan"
	if length := len(name); length > 10 {
		fmt.Println("Name is long")
	} else {
		fmt.Println("name is short")
	}

}
