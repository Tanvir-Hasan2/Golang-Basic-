package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file practice ")
	message := "File is ready!!"

	file, err := os.Create("./1-textFile.txt")
	checkNilError(err)

	defer file.Close()
	file.WriteString(message)

	// length, err := io.WriteString(file, message)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("lenght is : ", length)

	readFile("./1-textFile.txt")

}
func readFile(file string) {
	messageInBytes, err := ioutil.ReadFile(file)
	checkNilError(err)
	fmt.Println("data from text file is : ", string(messageInBytes))
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
