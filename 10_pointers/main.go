package main

import "fmt"

func main() {
	a := 5
	b := &a

	// Reading memory address
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

}
