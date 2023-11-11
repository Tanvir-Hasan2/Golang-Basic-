package main

import "fmt"

func main() {
	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)
}

//syntax
// func FunctionName(param1 type, param2 type) type {
// 	// code to be executed
// 	return output
// }
func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}
