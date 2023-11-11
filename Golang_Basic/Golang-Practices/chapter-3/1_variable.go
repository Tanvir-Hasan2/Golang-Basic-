package main

import "fmt"

func main() {
	// Creating a rune
	//runes are used to represent individual Unicode code points
	var r rune
	r = 'A' // You can assign a character to a rune using single quotes
	fmt.Printf("Rune: %c\n", r)

	// Iterating over a string and printing its runes
	str := "Hello, 世界" // Contains both ASCII and non-ASCII characters
	for i, runeValue := range str {
		fmt.Printf("Character at index %d: %c\n", i, runeValue)
	}
}
