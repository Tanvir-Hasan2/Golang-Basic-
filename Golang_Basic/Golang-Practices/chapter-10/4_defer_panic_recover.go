package main

import "fmt"

func deferfunction() {
	fmt.Println("defer function")
	if r := recover(); r != nil {
		fmt.Println("recovered panic :", r)
	}
}

func main() {
	defer deferfunction() // called before the function return
	fmt.Println("before panic")
	panic("something bad happend") // stop the execution after panic occurred
	fmt.Println("after panic")     //unreachable statement
}
