// Package data defines data models.
package data

import "fmt"

// Person represents properties of a person.
type Person struct {
	Name   string
	Age    int
	Gender string
}

// Describe prints a person's properties
func (p Person) Describe() {
	fmt.Printf("%v is %v years old, gender is %v.\n", p.Name, p.Age, p.Gender)
}

// SetName changes a person's name
func (p *Person) SetName(Name string) {
	p.Name = Name
}

// SetAge changes a person's age
func (p *Person) SetAge(Age int) {
	p.Age = Age
}

// SetGender changes a person's gender (if that's possible...)
func (p *Person) SetGender(Gender string) {
	p.Gender = Gender
}
