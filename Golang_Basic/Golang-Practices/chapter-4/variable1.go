package main

import "fmt"

func main() {
	var name string = "Anik Adnan"
	fmt.Println(name)

	name = "Adnan"
	fmt.Println(name)
	name = name + " " + "Biswas" // varobale can be concatenation
	fmt.Println(name)

	//equality check ==
	name1 := "anik" // : automaticaly type is assigned by the compiler
	name2 := "anik"
	fmt.Println("check string is equal or not :", name1 == name2) // true
}
