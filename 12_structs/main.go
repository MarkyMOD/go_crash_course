package main

import (
	"fmt"
	"strconv"
)

// Define Person Struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (Value Receiver)

func (p Person) greet() string {
	return "Hello my name is " + p.firstName + "" + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// HasBirthday Method (Pointer Receiver)

func (p *Person) hasBirthday() {
	p.age++
}

// Get Married (Pointer Receiver)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	} else {
		return
	}
}

func main() {
	// Init person using struct

	person1 := Person{firstName: "Samantha ", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative

	// person1 := Person{"Samantha", "Smith", "Boston", "f", 25}

	// fmt.Println(person1)
	// person1.age++
	// fmt.Println(person1.age)

	person1.hasBirthday()
	person1.hasBirthday()

	person1.getMarried("Williams")

	fmt.Println(person1.greet())
}
