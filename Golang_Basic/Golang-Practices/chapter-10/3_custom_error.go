package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	msg string
}

func (err CustomError) Error() string {
	return fmt.Sprintf("error massage: %s", err.msg)
}
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, CustomError{msg: "divide by zero"}
	}
	if a%b != 0 {
		// return 0, CustomError{msg: "not divisible"}
		return 0, errors.New("not divisible")
	}
	return a / b, nil
}
func main() {
	res, err := Divide(5, 0)
	if err != nil {
		// error unwrapping
		var cErr CustomError
		if errors.As(err, &cErr) {
			fmt.Println("custom error occurred")
			fmt.Printf("Error : %s\n", cErr.msg)
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(res)
}
