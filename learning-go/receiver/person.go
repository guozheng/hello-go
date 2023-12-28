package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func (p Person) hello() string {
	return fmt.Sprintf("Hello, my name is: %s %s, I am %d years old", p.firstName, p.lastName, p.age)
}

func (p *Person) celebrateBirthday() {
	fmt.Println("Celebrating my birthday, getting older...")
	p.age++
}

func main() {
	p := Person{
		firstName: "John",
		lastName: "Doe",
		city: "New York",
		gender: "Male",
		age: 30,
	}
	fmt.Println(p.hello())

	p.celebrateBirthday()
	fmt.Println(p.hello())
}