package main

import "fmt"

func main() {
	// Assigning an anonymous
	// function to a variable
	value := func() {
		fmt.Println("Welcome!")
	}
	value()
	fmt.Printf("%T", value) // Type: func()
}
