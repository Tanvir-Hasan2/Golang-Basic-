package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println(b)  // memory reference
	fmt.Println(*b) //value

	x := 0
	zero(&x)
	fmt.Println(&x, x)
	fmt.Println()

	ptr := new(int) //
	zero(ptr)
	fmt.Println(ptr, *ptr)

}
func zero(xptr *int) {
	*xptr = 10
	fmt.Println(xptr, *xptr)
}
