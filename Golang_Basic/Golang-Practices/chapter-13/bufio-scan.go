package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter bufio scan intput in console  ")
	scanner.Scan()
	str := scanner.Text()
	fmt.Println("text : ", str)

	// read from file
	fmt.Println("read from file bufio.text")

	f, err := os.Open("bufio_scan_input_read_write.text")
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
