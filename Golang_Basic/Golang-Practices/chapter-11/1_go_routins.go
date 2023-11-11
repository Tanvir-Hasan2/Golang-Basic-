package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 'a'; i <= 'z'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}
func printLetters() {
	for i := 1; i <= 20; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	go printNumbers()
	go printLetters()
	// time.Sleep(500 * time.Millisecond)
	// if main function finished all the go routines are finished
	fmt.Println("main function finished")
}
