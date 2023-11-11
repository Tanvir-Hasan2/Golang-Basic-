package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int32 = 11
	var i1 int64 = int64(i)

	var f float32 = 32
	var f2 float64 = float64(f)

	var s1 int = 65
	var s2 string = strconv.Itoa(s1)

	var a int = 65 //"A" ascii value
	var b string = string(a)

	fmt.Printf("%T %T", i, i1)
	fmt.Printf("\n%T %T", f, f2)

	fmt.Printf("\n %v %T", s2, s2)
	fmt.Printf("\n %v %T", b, b)
	
}
