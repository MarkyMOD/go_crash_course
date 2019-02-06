package main

import "fmt"

func main() {
	x := 5
	y := 3

	// If Else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else If
	color := "red"

	if color == "red" {
		fmt.Printf("color is red\n")
	} else if color == "blue" {
		fmt.Printf("color is blue\n")
	} else {
		fmt.Printf("color is not blue, or red\n")
	}

	// Switch
	switch color {
	case "red":
		fmt.Printf("color is red\n")
	case "blue":
		fmt.Printf("color is blue\n")
	default:
		fmt.Printf("color is not blue, or red\n")

	}
}
