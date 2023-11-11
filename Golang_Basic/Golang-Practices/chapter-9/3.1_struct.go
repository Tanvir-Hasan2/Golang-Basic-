package main

import "fmt"

type Person struct {
	name string
	age  int32
	sex  string
}

//refernce type
func (p *Person) changeAge() {
	p.age = p.age - 4
}

//value type
func (p Person) changeName() {
	p.name = "No Name" //change not happened
}
func main() {
	p := Person{name: "Anik Adnan", age: 23, sex: "unmarried"}
	fmt.Println(p.sex)
	p.changeName()
	fmt.Println(p.name)

}
