package main

import "fmt"

func main() {

	// Define Map
	// emails := make(map[string]string)

	// // Assign Key:Value Pairs
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare Map, and Add KV

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete From Map
	delete(emails, "Bob")

	fmt.Println(emails)

}
