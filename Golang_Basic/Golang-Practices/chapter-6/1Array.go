package main

import "fmt"

func main() {
	var a [5]int // by default all are zero valued

	a[2] = 100
	fmt.Println(a)
	fmt.Println(a[2])

	fmt.Print("enter five number : ")
	i := 0
	for i < 5 {
		fmt.Scanf("%d", &a[i])
		i += 1
	}
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += float64(a[i])

	}
	fmt.Println(total / float64(len(a)))

}
