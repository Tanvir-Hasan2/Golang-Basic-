package main

import (
	"fmt"
)

func main() {
	// var name map[string]int // compile time and runtime error
	name := make(map[string]int) // map os unordered collections of data
	name["anik"] = 25
	fmt.Println(name)
	fmt.Println(name["anik"])
	name["adnan"] = 30
	name["biswas"] = 35
	fmt.Println(name)
	delete(name, "biswas")
	fmt.Println("delete biswas:", name)

	value_age, ok := name["anik"] // not pesent return 0, false

	fmt.Println(value_age, ok) // ok true/false

}
