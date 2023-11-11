package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{5}
	r := Rectangle{width: 4, height: 8}

	fmt.Printf("Circle area: %v Perimeter : %v", c.area(), c.perimeter())
	fmt.Printf("\nrectangle area: %v Perimeter : %v", r.area(), r.perimeter())
}

type Shape interface {
	area() float64
	perimater() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}
