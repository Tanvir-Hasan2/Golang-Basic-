package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read from file
	fmt.Println("read from file bufio.text")

	f, err := os.Open("read.text")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// read from a file
	fileReader := bufio.NewScanner(f)
	for fileReader.Scan() {
		fmt.Println(fileReader.Text())
	}
	if err := fileReader.Err(); err != nil {
		fmt.Println(err)
	}

}
