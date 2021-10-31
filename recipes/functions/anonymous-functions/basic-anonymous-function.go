package main

import "fmt"

func main() {
	// Anonymous function
	func() {
		fmt.Println("Inside anonymous function")
	}() // calling with ()
}
