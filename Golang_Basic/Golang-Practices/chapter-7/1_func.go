package main

import "fmt"

// func FunctionName(param1 type, param2 type, param3 type) {
// 	// code to be executed
//   }

func main() {
	x := []float64{10.5, 403, 55, 7, 9}
	result := average(x)
	fmt.Println(result)

	a, b := returnMultipleValue()
	fmt.Println(a, b)

	r := returnNameType()
	fmt.Println(r)

}
func average(a []float64) float64 {
	total := 0.0
	for _, value := range a {
		total += value

	}
	return total / float64(len(a))
}
func returnMultipleValue() (int, int) {
	return 5, 10
}

func returnNameType() (result int) {
	result = 70
	return
}
