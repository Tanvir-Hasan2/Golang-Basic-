package main

import "fmt"

func main() {
	fmt.Println("Enter dgit")
	var digit int
	fmt.Scanf("%d ", &digit)

	// switch initialization; condition
	switch digit {
	case 0, 11: // multiple cases we can decalared
		fmt.Println("Zero , eleven")
		// fallthrough
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	case 7:
		fmt.Println("seven")
	case 8:
		fmt.Println("eight")
	case 9:
		fmt.Println("nine")
	default:
		fmt.Println("not a number")

	}

}
