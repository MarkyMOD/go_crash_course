package main

import "fmt"

func main() {
	//using var

	// var name = "Mark"
	// Shorthand for var
	name := "Mark"

	var age int32 = 31

	var isCool = true
	isCool = false

	monthNumber, word := 1, "word"

	fmt.Println(name, age, isCool, monthNumber, word)
	fmt.Printf("%T\n", age)

}
