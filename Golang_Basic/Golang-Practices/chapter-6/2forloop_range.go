package main

import "fmt"

func main() {
	array := [4]int{1, 2, 3, 4} // value initializea,donot need to type
	var x [3]int
	fmt.Println(array)
	fmt.Println("enter three number")
	for i := 0; i < len(x); i++ {
		fmt.Scanf("%d", &x[i])
	}

	var total float64 = 0
	for _, value := range x { //for index,value := range x
		total += float64(value)
		fmt.Println(value)
	}
	fmt.Print("Total value is : ", total)
}
