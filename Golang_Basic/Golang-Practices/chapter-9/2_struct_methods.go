package main

import (
	"fmt"
	"math"
)

func main() {
	var c Circle
	fmt.Println(c.x, c.y, c.r) // 0 0 0
	c1 := Circle{0, 0, 2}
	fmt.Println(circleArea(&c1))
	fmt.Println(c1.area()) // donot need to pass & operator

}

type Circle struct {
	x, y, r float64
}

//reference type
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
