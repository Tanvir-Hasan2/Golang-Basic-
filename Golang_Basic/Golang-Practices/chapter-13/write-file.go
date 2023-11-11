package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages ")
	scanner.Scan()
	message := scanner.Text()
	fmt.Println("text-messages : ", message)

	// create a file
	f, err := os.Create("write.text")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// write meassage in the file
	f.WriteString(message)

	// read from a file
	fileReader := bufio.NewScanner(f)
	for fileReader.Scan() {
		fmt.Println(fileReader.Text())
	}
	if err := fileReader.Err(); err != nil {
		fmt.Println(err)
	}

}
