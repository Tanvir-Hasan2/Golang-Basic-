package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("give intput ")
	scanner.Scan()
	message := scanner.Text()
	fmt.Println("text : ", message)

	// read from file
	fmt.Println("from file")
	f, err := os.Open("bufio.text")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// write in afile
	// f.WriteString(message)

}
