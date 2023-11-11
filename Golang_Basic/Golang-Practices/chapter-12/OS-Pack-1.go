package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new file.
	f, err := os.Create("my_file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Write some data to the file.
	data := []byte("My name is Anik Adnan")
	f.Write(data)

	fmt.Println("done ")
}
