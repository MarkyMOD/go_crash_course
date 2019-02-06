package main

import "fmt"

func main() {
	// Arrays

	// var fruitArr [2]string

	// // Assign Values

	// fruitArr[0] = "Apples"
	// fruitArr[1] = "Orange"

	// Declare and Assign Array

	fruitArr := [2]string{"Apple", "Orange"}

	// Slices

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr[0], fruitArr, "Slice:", fruitSlice)
}
