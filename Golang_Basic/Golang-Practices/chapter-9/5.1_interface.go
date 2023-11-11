package main

import "fmt"

type I interface {
}

func displayType(i I) {
	switch i.(type) {
	case int:
		fmt.Println("integar type")
	case float64:
		fmt.Println("float type")
	case bool:
		fmt.Println("bollean type")
	default:
		fmt.Println("other type")

	}
}
func main() {

	displayType("anik")
	displayType(true)
	displayType(3.54)
}
