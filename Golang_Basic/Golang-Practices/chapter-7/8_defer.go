package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	defer second() //defer move to the callto second the end of the function
	first()
	first()
	// f, _ := os.Open(filename)
	// defer f.Close()
}
