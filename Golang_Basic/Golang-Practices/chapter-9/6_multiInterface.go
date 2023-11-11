package main

import "fmt"

func main() {
	var x float64
	fmt.Println(x)

}

type MultiShape struct {
	shapes []Shape
}
type Shape interface {
	area() float64
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
