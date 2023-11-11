package main

import "fmt"

const (
	// group variable incrimented alwayas
	x = iota
	y = iota
	z = iota
)
const (
	i = iota //0
	_        // 1
	_        // 2
	// j = iota
	j //same as before
)

func main() {
	const con int = 10
	const c = iota
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(i)
	fmt.Println(j)

}
