package main // package declaration

import "fmt" // go doc fmt

// this is comment

func main() {
	fmt.Println("hello world")

	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.56+1.0)

	fmt.Println(len("hello wrold"))
	fmt.Println("Hello World"[1]) // 101 = "e"
	fmt.Println("hello" + "world")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
