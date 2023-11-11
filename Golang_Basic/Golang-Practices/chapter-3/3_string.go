package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = " aanik"
	var s2 string = " adNaN"
	var str string = "aniK@ aDnAn"
	var mail string = "  anikadnanict17@gmail.com@@"
	fmt.Println(strings.Trim(mail, " ")) //trim
	fmt.Println(strings.TrimSuffix(mail, "@"))
	fmt.Println(strings.ToLower(str))
	// fmt.Println(strings.ToLowerSpecial("@",str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToTitle("anik adnan"))

	fmt.Println(strings.Count(s1, "a"))

	fmt.Println(strings.Contains(s2, "nan"))    //case not ingore
	fmt.Println(strings.ContainsAny(s2, "nan")) // ingore case
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s2))

}
