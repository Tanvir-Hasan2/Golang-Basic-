package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	result := float64(a / b)
	return result, nil
}
func main() {
	div, err := Divide(7, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println(div)
}
