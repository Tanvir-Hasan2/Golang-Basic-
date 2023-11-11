package main

import (
	"fmt"
	"math"
)

func main() {

	// var c Circle
	// c :=new(Circle) //(*Circle)

	c1 := Circle{x: 0, y: 0, r: 5}
	c2 := Circle{0, 0, 3}

	fmt.Println(c1, c2)
	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(c2.x, c2.y, c2.r)
	fmt.Println(circleArea(c1))   // copy of c1
	fmt.Println(circleArea2(&c2)) // origial location

	v := &c2 //refer to same memory location
	fmt.Println(v)

}

type Circle struct {
	x, y, r float64
}

// value type
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// reference type
func circleArea2(c *Circle) float64 {
	c.r = 7
	return math.Pi * c.r * c.r
}
