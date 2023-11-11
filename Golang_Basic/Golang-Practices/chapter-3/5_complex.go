package main

import "fmt"

func main() {
	var a float64 = 2
	var b float64 = 4
	var c complex128 = complex(a, b)
	fmt.Println(real(c), imag(c))
	fmt.Println(c)
}
