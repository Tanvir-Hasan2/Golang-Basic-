package main

import "fmt"

func main() {
	a := new(Android)
	// a.person.name = "Anik"
	// fmt.Println(a.person.name)

	// a.person.Talk()
	a.Talk()
}

type Person struct { // strict's fields usually has a relationship
	name string
}
type Android struct {
	// person Person
	Person // donot give a reference name for embaded //android has a person
	model  string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.name)
}
