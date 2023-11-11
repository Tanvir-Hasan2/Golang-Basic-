package main

import "fmt"

const PI = 3.14 //Untyped constant
//const block foe multiple declaration
const (
	A int = 0
	S     = "$"
)

func main() {
	const name string = "No name" // typed constant
	// cannot reinitalizea ,it's a compiler time error
	fmt.Println(name)
}
