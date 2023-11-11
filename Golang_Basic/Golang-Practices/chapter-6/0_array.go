package main

import "fmt"

func main() {
	var aa []int = []int{1, 2}
	fmt.Println(aa)
	var a = [5]int64{1, 2, 3, 4, 5} //length is define
	fmt.Println(a)
	var ar = [...]int64{1, 2, 3} // length is infered
	ar[2] = 50
	fmt.Println(ar, ar[2])

	// x:=[3]int64{1,2,3} //length define
	// y:=[...]int64{1,2} // length is infered

	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1, arr2, arr3)

	arr := [5]int{1: 10, 2: 40} // specific  initialization

	fmt.Println(arr)
	fmt.Println(len(arr)) // length
}
