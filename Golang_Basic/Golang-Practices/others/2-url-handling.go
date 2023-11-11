package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://myInfo.com:300/profile?AnikAdnan-28=reactjs&ID=ghbscd2395fcs"

func main() {
	fmt.Println("Handling url in golang")

	//parsing

	result, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("scheme :", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("path:", result.Path)
	fmt.Println("port:", result.Port())
	fmt.Println("rawQuary:", result.RawQuery)
	fmt.Println("user:", result.User)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)
	fmt.Println(qparams["AnikAdnan-28"])

	for _, val := range qparams {
		fmt.Println("param is :", val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "myInfo.com",
		Path:    "/profile",
		RawPath: "user=AnikAnik",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)

}
