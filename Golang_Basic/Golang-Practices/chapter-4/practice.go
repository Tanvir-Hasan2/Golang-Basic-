package main

import "fmt"

func main() {
	var a int
	var b, c int = 6, 7 // perline one type, if using  var keyword
	var s1 string       // ""
	var age int         // 0
	fmt.Println(a, b, c)
	fmt.Println(s1, age)
	// veriable decalaration in a  block
	var (
		x int
		y int    = 10
		s string = "Hello"
	)

	fmt.Println(x, y, s)
}
