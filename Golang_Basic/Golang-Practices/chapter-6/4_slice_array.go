package main

import "fmt"

func main() {
	// slice_name := make([]type, length, capacity)

	arr := []float64{1, 2, 3, 4}
	fmt.Println(arr)

	x := arr[0:3]
	fmt.Println(x)

	// slice_name = append(slice_name, element1, element2, ...)
	y := append(x, 10, 12, 13, 1, 23, 5)
	fmt.Println(y)

	y = append(y, x...)

	fmt.Println(y)
	fmt.Println(len(y), cap(y))

	// copy(dest, src)
	copy(x, y) //memory efficency
	fmt.Println(x)

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	// slice3 = append(slice1, slice2...)
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}
