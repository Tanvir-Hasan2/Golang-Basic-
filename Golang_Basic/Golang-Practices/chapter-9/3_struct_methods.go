package main

import (
	"fmt"
	"math"
)

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.x1, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y1)
	return l * w
}
func distance(x1, x2, y1, y2 float64) float64 {
	x := x1 - x2
	y := y1 - y2
	return math.Sqrt(x*x + y*y)
}
