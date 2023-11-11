package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://127.0.0.1:5500/Golang-Practices/others/1-info.html"

func main() {
	fmt.Println("web request ")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	fmt.Println("response is of type : %T\n", response)
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	content := string(dataBytes)
	fmt.Println(content)

}
