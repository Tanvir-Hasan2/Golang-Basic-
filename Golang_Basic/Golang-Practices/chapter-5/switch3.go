package main

import (
	"fmt"
)

func main() {
	var x interface{} = "this is string"

	switch x.(type) {
	case int:
		fmt.Println("int type")
	case float32:
		fmt.Println("float type")
	case string:
		fmt.Println("string type")
	default:
		fmt.Println("complex type")
	}
}
